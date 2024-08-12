package service

import (
	"btb-service/model"
	"btb-service/pkg"
	"btb-service/repository"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func InsertUser(ctx *fiber.Ctx) error {
	var payload = &model.UserInsertPayload{}
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

	data, err := repository.GetUserByUsernameOrEmail(payload.Username, payload.Username)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":      "USER.USERQUERY.EXCEPTION",
			"message":    "Failed to validate user data!",
			"stacktrace": err.Error(),
		})
	}

	if len(data) != 0 {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error":      "USER.USEREXIST.EXCEPTION",
			"message":    "Sorry, User is already exist in database!",
			"stacktrace": "User is already exist",
		})
	}

	payload.Password = pkg.HashPasswordBCrypt(payload.Password)
	payload.IsActive = true
	if err := repository.SaveUser(*payload); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":      "USER.REGISTERUSER.EXCEPTION",
			"message":    "Failed to USER!",
			"stacktrace": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Success USER!",
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
