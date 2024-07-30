package service

import (
	"btb-service/model"
	"btb-service/pkg"
	"btb-service/repository"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func SignUp(ctx *fiber.Ctx) error {
	var payload = &model.UserInsertPayload{}
	if err := ctx.BodyParser(payload); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":      "SIGNUP.INVALIDPAYLOAD.EXCEPTION",
			"message":    "Sorry, System can't parse your data! Please Recheck!",
			"stacktrace": err.Error(),
		})
	}

	var goValidator = validator.New()
	if err := goValidator.Struct(payload); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":      "SIGNUP.INVALIDPAYLOAD.EXCEPTION",
			"message":    "Parameter is required!",
			"stacktrace": err.Error(),
		})
	}

	data, err := repository.GetUserByUsernameOrEmail(payload.Username, payload.Username)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":      "SIGNUP.USERQUERY.EXCEPTION",
			"message":    "Failed to validate user data!",
			"stacktrace": err.Error(),
		})
	}

	if len(data) != 0 {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error":      "SIGNUP.USEREXIST.EXCEPTION",
			"message":    "Sorry, User is already exist in database!",
			"stacktrace": "User is already exist",
		})
	}

	payload.Password = pkg.HashPasswordBCrypt(payload.Password)
	payload.IsActive = true
	if err := repository.SaveUser(*payload); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":      "SIGNUP.REGISTERUSER.EXCEPTION",
			"message":    "Failed to signup!",
			"stacktrace": err.Error(),
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
			"error":      "SIGNIN.INVALIDPAYLOAD.EXCEPTION",
			"message":    "Sorry, System can't parse your data! Please Recheck!",
			"stacktrace": err.Error(),
		})
	}

	var goValidator = validator.New()
	if err := goValidator.Struct(payload); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":      "SIGNIN.INVALIDPAYLOAD.EXCEPTION",
			"message":    "Parameter is required!",
			"stacktrace": err.Error(),
		})
	}

	data, err := repository.GetUserByUsernameOrEmail(payload.Username, payload.Username)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":      "SIGNIN.USERQUERY.EXCEPTION",
			"message":    "Failed to validate user data!",
			"stacktrace": err.Error(),
		})
	}

	if len(data) == 0 {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":      "SIGNIN.USERNOTEXIST.EXCEPTION",
			"message":    "Sorry, User is not exist in database!",
			"stacktrace": "User is not exist",
		})
	}

	if !pkg.ComparePasswordBCrypt(payload.Password, data[0]["password"].(string)) {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":      "SIGNIN.INVALIDPASSWORD.EXCEPTION",
			"message":    "Password not match!",
			"stacktrace": "Incorrect Password!",
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
			"error":      "SIGNIN.TOKEN.EXCEPTION",
			"message":    "Failed generate token!",
			"stacktrace": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Success Signin!",
		"data": fiber.Map{
			"type":       "Bearer",
			"token":      jwtToken,
			"expiredate": expiredDate,
		},
	})
}

func Validate(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Success Validate Token!",
		"data":    ctx.Locals("jwtauth"),
	})
}

func GetAllUser(ctx *fiber.Ctx) error {
	var payload = &map[string]interface{}{}
	if err := ctx.BodyParser(payload); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":      "USER.INVALIDPAYLOAD.EXCEPTION",
			"message":    "Sorry, System can't parse your data! Please Recheck!",
			"stacktrace": err.Error(),
		})
	}

	data, err := repository.GetAllUser(*payload)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":      "USER.QUERY.EXCEPTION",
			"message":    "Failed to get User data!",
			"stacktrace": err.Error(),
		})
	}

	if len(data) == 0 {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":      "USER.NOTEXIST.EXCEPTION",
			"message":    "Sorry, USER Data isn't exist!",
			"stacktrace": "User data is not exist",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Success get User Data!",
		"data":    data,
	})
}

func UpdateUser(ctx *fiber.Ctx) error {
	var payload = &model.UserUpdatePayload{}
	if err := ctx.BodyParser(payload); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":      "USER.INVALIDPAYLOAD.EXCEPTION",
			"message":    "Sorry, System can't parse your data! Please Recheck!",
			"stacktrace": err.Error(),
		})
	}

	var goValidator = validator.New()
	if err := goValidator.Struct(payload); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":      "USER.INVALIDPAYLOAD.EXCEPTION",
			"message":    "Parameter is required!",
			"stacktrace": err.Error(),
		})
	}

	err := repository.UpdateUser(*payload)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":      "USER.SUBMIT.EXCEPTION",
			"message":    "Failed to update user!",
			"stacktrace": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Success submit data!",
	})
}
