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
			"code":    "CAREER.APPLICANTQUERY.EXCEPTION",
			"message": "Failed to get applicant data!",
			"error":   err.Error(),
		})
	}

	if len(data) == 0 {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"code":    "CAREER.APPLICANTNOTEXIST.EXCEPTION",
			"message": "Sorry, Applicant Data isn't exist!",
			"error":   "Applicant data is not exist",
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
			"code":    "CAREER.APPLICANTINVALIDPAYLOAD.EXCEPTION",
			"message": "Sorry, System can't parse your data! Please Recheck!",
			"error":   err.Error(),
		})
	}

	var goValidator = validator.New()
	if err := goValidator.Struct(payload); err != nil {

		var errorMessage string
		for _, err := range err.(validator.ValidationErrors) {
			fieldName := err.StructField()
			switch err.Tag() {
			case "required":
				errorMessage += fieldName + " is required.<br/>"
			case "email":
				errorMessage += fieldName + " must be a valid email address.<br/>"
			case "e164":
				errorMessage += fieldName + " must be a valid Phone no<br/>"
			default:
				errorMessage += fieldName + " is invalid.<br/>"
			}
		}

		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"code":    "CAREER.INVALIDPAYLOAD.EXCEPTION",
			"message": errorMessage,
			"error":   errorMessage,
		})
	}

	err := repository.ApplyCareer(*payload)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"code":    "CAREER.APPLICANTSUBMIT.EXCEPTION",
			"message": "Failed to apply Career!",
			"error":   err.Error(),
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Success submit data!",
	})
}

func UpsertCareer(ctx *fiber.Ctx) error {
	var payload = &model.CareerUpsertPayload{}
	if err := ctx.BodyParser(payload); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"code":    "CAREER.INVALIDPAYLOAD.EXCEPTION",
			"message": "Sorry, System can't parse your data! Please Recheck!",
			"error":   err.Error(),
		})
	}

	var goValidator = validator.New()
	if err := goValidator.Struct(payload); err != nil {

		var errorMessage string
		for _, err := range err.(validator.ValidationErrors) {
			fieldName := err.StructField()
			switch err.Tag() {
			case "required":
				errorMessage += fieldName + " is required.<br/>"
			case "email":
				errorMessage += fieldName + " must be a valid email address.<br/>"
			case "e164":
				errorMessage += fieldName + " must be a valid Phone no<br/>"
			default:
				errorMessage += fieldName + " is invalid.<br/>"
			}
		}

		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"code":    "CAREER.INVALIDPAYLOAD.EXCEPTION",
			"message": errorMessage,
			"error":   errorMessage,
		})
	}

	err := repository.UpsertCareer(*payload)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"code":    "CAREER.SUBMIT.EXCEPTION",
			"message": "Failed to submit Career!",
			"error":   err.Error(),
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Success submit data!",
	})
}

func GetActiveCareer(ctx *fiber.Ctx) error {
	var payload = &map[string]interface{}{}
	if err := ctx.BodyParser(payload); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"code":    "CAREER.INVALIDPAYLOAD.EXCEPTION",
			"message": "Sorry, System can't parse your data! Please Recheck!",
			"error":   err.Error(),
		})
	}

	data, err := repository.GetActiveCareer(*payload)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"code":    "CAREER.QUERY.EXCEPTION",
			"message": "Failed to get career data!",
			"error":   err.Error(),
		})
	}

	if len(data) == 0 {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"code":    "CAREER.NOTEXIST.EXCEPTION",
			"message": "Sorry, Career Data isn't exist!",
			"error":   "Career data is not exist",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Success get Career Data!",
		"data":    data,
	})
}

func GetAllCareer(ctx *fiber.Ctx) error {
	var payload = &map[string]interface{}{}
	if err := ctx.BodyParser(payload); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"code":    "CAREER.INVALIDPAYLOAD.EXCEPTION",
			"message": "Sorry, System can't parse your data! Please Recheck!",
			"error":   err.Error(),
		})
	}

	data, err := repository.GetAllCareer(*payload)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"code":    "CAREER.QUERY.EXCEPTION",
			"message": "Failed to get career data!",
			"error":   err.Error(),
		})
	}

	if len(data) == 0 {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"code":    "CAREER.NOTEXIST.EXCEPTION",
			"message": "Sorry, Career Data isn't exist!",
			"error":   "Career data is not exist",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Success get Career Data!",
		"data":    data,
	})
}

func DeleteCareer(ctx *fiber.Ctx) error {
	var payload = &model.CareerDeletePayload{}
	if err := ctx.BodyParser(payload); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"code":    "CAREER.INVALIDPAYLOAD.EXCEPTION",
			"message": "Sorry, System can't parse your data! Please Recheck!",
			"error":   err.Error(),
		})
	}

	var goValidator = validator.New()
	if err := goValidator.Struct(payload); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"code":    "CAREER.INVALIDPAYLOAD.EXCEPTION",
			"message": "Parameter is required!",
			"error":   err.Error(),
		})
	}

	if err := repository.DeleteCareer(*payload); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"code":    "CAREER.QUERY.EXCEPTION",
			"message": "Failed to delete career data!",
			"error":   err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Success delete career data",
	})
}
