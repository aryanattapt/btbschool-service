package service

import (
	"btb-service/model"
	"btb-service/pkg"
	"btb-service/repository"
	"fmt"
	"log"

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

	if payload.Status == "send" {
		log.Println("validate form online registration")
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
				"error":      "STUDENTREG.INVALIDPAYLOAD.EXCEPTION",
				"message":    errorMessage,
				"stacktrace": errorMessage,
			})
		}
	}

	registrationCode, err := repository.SubmitDataStudentRegistration(*payload)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":      "STUDENTREG.SUBMIT.EXCEPTION",
			"message":    "Failed to submit student data!",
			"stacktrace": err.Error(),
		})
	}

	var message string
	if pkg.IsEmptyString(payload.RegistrationCode) {
		message = fmt.Sprintf(`
			<html>
			<head>
				<style>
					body { font-family: Arial, sans-serif; }
					.container { margin: 20px; }
					.footer { margin-top: 20px; font-size: 0.9em; }
				</style>
			</head>
			<body>
				<div class="container">
					<p>Dear Mr/Mrs,</p>
					<p>Thank You For submit your data. Your registration code is <strong>%s</strong>.</p>
					<p>Thank you,</p>
					<p>Bina Tunas Bangsa School Admission</p>
				</div>
			</body>
			</html>
		`, registrationCode)
	} else {
		message = fmt.Sprintf(`
			<html>
			<head>
				<style>
					body { font-family: Arial, sans-serif; }
					.container { margin: 20px; }
					.footer { margin-top: 20px; font-size: 0.9em; }
				</style>
			</head>
			<body>
				<div class="container">
					<p>Dear Mr/Mrs,</p>
					<p>You're currently updated registration data for registration code <strong>%s</strong>.</p>
					<p>Thank you,</p>
					<p>Bina Tunas Bangsa School Admission</p>
				</div>
			</body>
			</html>
		`, payload.RegistrationCode)
	}

	var mailPayload pkg.MailPayload = pkg.MailPayload{
		To:      []string{payload.Email},
		Cc:      []string{},
		Subject: "Online Registration Notification",
		Message: message,
	}

	err = mailPayload.SendMail()
	if err != nil {
		log.Println(err)
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
