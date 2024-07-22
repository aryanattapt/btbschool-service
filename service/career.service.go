package service

import (
	"btb-service/model"
	"btb-service/repository"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func GetCareerApplicantData(ctx *fiber.Ctx) error {
	data, err := repository.GetCareerApplicantData()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":      "CAREER.APPLICANTQUERY.EXCEPTION",
			"message":    "Failed to get applicant data!",
			"stacktrace": err.Error(),
		})
	}

	if len(data) == 0 {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":      "CAREER.APPLICANTNOTEXIST.EXCEPTION",
			"message":    "Sorry, Applicant Data isn't exist!",
			"stacktrace": "Applicant data is not exist",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Success get Applicant Data!",
		"data":    data,
	})
}

func ApplyCareer(ctx *fiber.Ctx) error {
	var payload = &model.CareerApplyInsertPayload{}
	if err := ctx.BodyParser(payload); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":      "CAREER.APPLICANTINVALIDPAYLOAD.EXCEPTION",
			"message":    "Sorry, System can't parse your data! Please Recheck!",
			"stacktrace": err.Error(),
		})
	}

	var goValidator = validator.New()
	if err := goValidator.Struct(payload); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":      "CAREER.APPLICANTINVALIDPAYLOAD.EXCEPTION",
			"message":    "Parameter is required!",
			"stacktrace": err.Error(),
		})
	}

	err := repository.ApplyCareer(*payload)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":      "CAREER.APPLICANTSUBMIT.EXCEPTION",
			"message":    "Failed to apply Career!",
			"stacktrace": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Success submit data!",
	})
}
