package service

import (
	"btb-service/model"
	"btb-service/repository"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func GetDraftStudentRegistrationData(ctx *fiber.Ctx) error {
	var payload = &model.DraftStudentRegistrationData{}
	if err := ctx.BodyParser(payload); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":      "STUDENTREG.INVALIDPAYLOAD.EXCEPTION",
			"message":    "Sorry, System can't parse your data! Please Recheck!",
			"stacktrace": err.Error(),
		})
	}

	data, err := repository.GetDraftStudentRegistrationData(*payload)
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

func GetStudentRegistrationOutstandingData(ctx *fiber.Ctx) error {
	data, err := repository.GetStudentRegistrationOutstandingData()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":      "STUDENTREG.QUERY.EXCEPTION",
			"message":    "Failed to get outstanding data!",
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
		"message": "Success get student data!",
		"data":    data,
	})
}

func ApproveStudentRegistrationOutstandingData(ctx *fiber.Ctx) error {
	var payload = &map[string]interface{}{}
	if err := ctx.BodyParser(payload); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":      "STUDENTREG.INVALIDPAYLOAD.EXCEPTION",
			"message":    "Sorry, System can't parse your data! Please Recheck!",
			"stacktrace": err.Error(),
		})
	}

	err := repository.ApproveStudentRegistrationOutstandingData(ctx.Locals("userauth"), *payload)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":      "STUDENTREG.SUBMIT.EXCEPTION",
			"message":    "Failed to submit data!",
			"stacktrace": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Success submit data!",
	})
}

func GetAllStudentRegistrationAuthData(ctx *fiber.Ctx) error {
	data, err := repository.GetAllStudentRegistrationAuthData(ctx.Locals("userauth"))
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":      "STUDENTREG.QUERY.EXCEPTION",
			"message":    "Failed to get all student data!",
			"stacktrace": err.Error(),
		})
	}

	if len(data) == 0 {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":      "STUDENTREG.NOTEXIST.EXCEPTION",
			"message":    "Sorry, Student Registration Data isn't exist!",
			"stacktrace": "Student Registration data is not exist",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Success get student data!",
		"data":    data,
	})
}

func GetStudentRegistrationDetailData(ctx *fiber.Ctx) error {
	var payload = &map[string]interface{}{}
	if err := ctx.BodyParser(payload); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":      "STUDENTREG.INVALIDPAYLOAD.EXCEPTION",
			"message":    "Sorry, System can't parse your data! Please Recheck!",
			"stacktrace": err.Error(),
		})
	}

	data, err := repository.GetStudentRegistrationDetailData(*payload)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":      "STUDENTREG.QUERY.EXCEPTION",
			"message":    "Failed to get student data!",
			"stacktrace": err.Error(),
		})
	}

	if len(data) == 0 {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":      "STUDENTREG.NOTEXIST.EXCEPTION",
			"message":    "Sorry, Student Registration Data isn't exist!",
			"stacktrace": "Student Registration data is not exist",
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Success get student data!",
		"data":    data,
	})
}
