package service

import (
	"btb-service/model"
	"btb-service/pkg"
	"btb-service/repository"
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func SignUp(ctx *fiber.Ctx) error {
	var payload = &model.UserInsertPayload{}
	if err := ctx.BodyParser(payload); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"code":    "SIGNUP.INVALIDPAYLOAD.EXCEPTION",
			"message": "Sorry, System can't parse your data! Please Recheck!",
			"error":   err.Error(),
		})
	}

	var goValidator = validator.New()
	if err := goValidator.Struct(payload); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"code":    "SIGNUP.INVALIDPAYLOAD.EXCEPTION",
			"message": "Parameter is required!",
			"error":   err.Error(),
		})
	}

	data, err := repository.GetUserByUsernameOrEmail(payload.Username, payload.Username)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"code":    "SIGNUP.USERQUERY.EXCEPTION",
			"message": "Failed to validate user data!",
			"error":   err.Error(),
		})
	}

	if len(data) != 0 {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"code":    "SIGNUP.USEREXIST.EXCEPTION",
			"message": "Sorry, User is already exist in database!",
			"error":   "User is already exist",
		})
	}

	payload.Password = pkg.HashPasswordBCrypt(payload.Password)
	payload.IsActive = true
	if err := repository.SaveUser(*payload); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"code":    "SIGNUP.REGISTERUSER.EXCEPTION",
			"message": "Failed to signup!",
			"error":   err.Error(),
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Success Signup!",
	})
}

func SignIn(ctx *fiber.Ctx) error {
	var payload = &model.SignInPayload{}
	if err := ctx.BodyParser(payload); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"code":    "SIGNIN.INVALIDPAYLOAD.EXCEPTION",
			"message": "Sorry, System can't parse your data! Please Recheck!",
			"error":   err.Error(),
		})
	}

	var goValidator = validator.New()
	if err := goValidator.Struct(payload); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"code":    "SIGNIN.INVALIDPAYLOAD.EXCEPTION",
			"message": "Parameter is required!",
			"error":   err.Error(),
		})
	}

	data, err := repository.GetUserByUsernameOrEmail(payload.Username, payload.Username)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"code":    "SIGNIN.USERQUERY.EXCEPTION",
			"message": "Failed to validate user data!",
			"error":   err.Error(),
		})
	}

	if len(data) == 0 {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"code":    "SIGNIN.USERNOTEXIST.EXCEPTION",
			"message": "Sorry, User is not exist in database!",
			"error":   "User is not exist",
		})
	}

	if !pkg.ComparePasswordBCrypt(payload.Password, data[0]["password"].(string)) {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"code":    "SIGNIN.INVALIDPASSWORD.EXCEPTION",
			"message": "Password not match!",
			"error":   "Incorrect Password!",
		})
	}

	var jwtPayload = JWTPayload{
		StandardClaims: jwt.StandardClaims{
			Audience: pkg.EncodeBase64(data[0]["_id"].(primitive.ObjectID).Hex()),
			Id:       pkg.HashSHA1(pkg.GenerateCurrentTimeStamp()),
			Issuer:   pkg.HashMD5("BTB_SCHOOL_SYSTEM"),
			Subject:  "BTB_AUTH",
		},
		SessionID: pkg.HashMD5(pkg.GenerateRandomNumber(32)),
	}

	jwtToken, err, expiredDate := jwtPayload.CreateJWTToken()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"code":    "SIGNIN.TOKEN.EXCEPTION",
			"message": "Failed generate token!",
			"error":   err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Success Signin!",
		"data": fiber.Map{
			"type":       "Bearer",
			"token":      jwtToken,
			"expiredate": expiredDate * 1000,
		},
	})
}

func Validate(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Success Validate Token!",
		"data":    ctx.Locals("jwtauth"),
	})
}

func CheckPermission(ctx *fiber.Ctx) error {
	var payload = &model.UserCheckPermissionPayload{}
	if err := ctx.BodyParser(payload); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"code":    "CHECKPERMISSION.INVALIDPAYLOAD.EXCEPTION",
			"message": "Sorry, System can't parse your data! Please Recheck!",
			"error":   err.Error(),
		})
	}

	var goValidator = validator.New()
	if err := goValidator.Struct(payload); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"code":    "CHECKPERMISSION.INVALIDPAYLOAD.EXCEPTION",
			"message": "Parameter is required!",
			"error":   err.Error(),
		})
	}

	userauth := ctx.Locals("userauth").(map[string]interface{})
	var userid string
	if objID, ok := userauth["_id"].(primitive.ObjectID); ok {
		userid = objID.Hex()
	} else {
		log.Println("user id not ok")
	}

	err := repository.CheckPermission(userid, payload.Permission)
	if err != nil && err.Error() == "unauthorized" {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"code":    "CHECKPERMISSION.UNAUTHORIZED",
			"message": "Unauthorized Access",
		})
	} else if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"code":    "CHECKPERMISSION.EXCEPTION",
			"message": "Failed get data",
			"error":   err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Authorized!",
	})
}
