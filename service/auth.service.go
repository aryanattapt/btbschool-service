package service

import (
	"btb-service/model"
	"btb-service/pkg"
	"btb-service/repository"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func SignUp(ctx *fiber.Ctx) error {
	var payload = &model.UserInsertPayload{}
	if err := ctx.BodyParser(payload); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Sorry, System can't parse your data! Please Recheck!",
			"error":   "INVALID_PAYLOAD",
		})
	}

	var goValidator = validator.New()
	if err := goValidator.Struct(payload); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
			"error":   "INVALID_PAYLOAD",
		})
	}

	data, err := repository.GetUserByUsernameOrEmail(payload.Username, payload.Username)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
			"error":   "QUERY_EXCEPTION",
		})
	}

	if len(data) != 0 {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Sorry, User is already exist in database!",
			"error":   "INVALID_AUTH",
		})
	}

	payload.Password = pkg.HashPasswordBCrypt(payload.Password)
	payload.IsActive = true
	if err := repository.SaveUser(*payload); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
			"error":   "QUERY_EXCEPTION",
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
			"message": "Sorry, System can't parse your data! Please Recheck!",
			"error":   "INVALID_PAYLOAD",
		})
	}

	var goValidator = validator.New()
	if err := goValidator.Struct(payload); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
			"error":   "INVALID_PAYLOAD",
		})
	}

	data, err := repository.GetUserByUsernameOrEmail(payload.Username, payload.Username)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
			"error":   "QUERY_EXCEPTION",
		})
	}

	if len(data) == 0 {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Sorry, User is not exist in database!",
			"error":   "INVALID_AUTH",
		})
	}

	if !pkg.ComparePasswordBCrypt(payload.Password, data[0]["password"].(string)) {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Sorry, Invalid Password!",
			"error":   "INVALID_AUTH",
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
			"message": err.Error(),
			"error":   "INVALID_AUTH",
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
	result, err := ValidateJWTToken(strings.Split(string(ctx.Request().Header.Peek("Authorization")), " ")[1])
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": err.Error(),
			"error":   "INVALID_AUTH",
		})
	}

	userDataList, err := repository.GetUserById(pkg.DecodeBase64(result["aud"].(string)))
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
			"error":   "QUERY_EXCEPTION",
		})
	}

	if len(userDataList) == 0 {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Sorry, User is not exist in database!",
			"error":   "INVALID_AUTH",
		})
	}

	var userData map[string]interface{} = userDataList[0]
	delete(userData, "_id")
	delete(userData, "password")
	delete(userData, "isactive")

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Success Validate Token!",
		"data":    userData,
	})
}
