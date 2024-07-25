package service

import (
	"btb-service/model"
	"btb-service/repository"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func GetStudentRegistrationOutstandingData(ctx *fiber.Ctx) error {
	var payload = &model.StudentRegistrationOutstandingDataPayload{}
	if err := ctx.BodyParser(payload); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":      "STUDENTREG.INVALIDPAYLOAD.EXCEPTION",
			"message":    "Sorry, System can't parse your data! Please Recheck!",
			"stacktrace": err.Error(),
		})
	}

	data, err := repository.GetStudentRegistrationOutstandingData(*payload)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":      "STUDENTREG.QUERY.EXCEPTION",
			"message":    "Failed to get Student Registration Outstanding data!",
			"stacktrace": err.Error(),
		})
	}

	if len(data) == 0 {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":      "STUDENTREG.NOTEXIST.EXCEPTION",
			"message":    "Sorry, Student Registration Outstanding Data isn't exist!",
			"stacktrace": "Student Registration Outstanding data is not exist",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Success get Student Registration Outstanding data!",
		"data":    data,
	})
}

func SubmitDataStudentRegistration(ctx *fiber.Ctx) error {
	var payload = &model.StudentRegistrationInsertPayload{}
	if err := ctx.BodyParser(payload); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":      "STUDENTREG.INVALIDPAYLOAD.EXCEPTION",
			"message":    "Sorry, System can't parse your data! Please Recheck!",
			"stacktrace": err.Error(),
		})
	}

	var goValidator = validator.New()
	if err := goValidator.Struct(payload); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":      "STUDENTREG.INVALIDPAYLOAD.EXCEPTION",
			"message":    "Parameter is required!",
			"stacktrace": err.Error(),
		})
	}

	registrationCode, err := repository.SubmitDataStudentRegistration(*payload)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":      "STUDENTREG.SUBMIT.EXCEPTION",
			"message":    "Failed to submit student data!",
			"stacktrace": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Success submit student data!",
		"data": fiber.Map{
			"registrationCode": registrationCode,
		},
	})
}
