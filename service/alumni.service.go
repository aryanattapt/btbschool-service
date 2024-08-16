package service

import (
	"btb-service/model"
	"btb-service/pkg"
	"btb-service/repository"
	"fmt"
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetAlumni(ctx *fiber.Ctx) error {
	var payload = &map[string]interface{}{}
	if err := ctx.BodyParser(payload); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":      "ALUMNI.INVALIDPAYLOAD.EXCEPTION",
			"message":    "Sorry, System can't parse your data! Please Recheck!",
			"stacktrace": err.Error(),
		})
	}

	data, err := repository.GetAlumni(*payload)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":      "ALUMNI.ALUMNIQUERY.EXCEPTION",
			"message":    "Failed to get alumni data!",
			"stacktrace": err.Error(),
		})
	}

	if len(data) == 0 {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":      "ALUMNI.NOTEXIST.EXCEPTION",
			"message":    "Sorry, Alumni Data isn't exist!",
			"stacktrace": "Alumni data is not exist",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Success get alumni!",
		"data":    data,
	})
}

func SubmitAlumni(ctx *fiber.Ctx) error {
	var payload = &model.AlumniInsertPayload{}
	if err := ctx.BodyParser(payload); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":      "ALUMNI.INVALIDPAYLOAD.EXCEPTION",
			"message":    "Sorry, System can't parse your data! Please Recheck!",
			"stacktrace": err.Error(),
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
			"error":      "ALUMNI.INVALIDPAYLOAD.EXCEPTION",
			"message":    errorMessage,
			"stacktrace": errorMessage,
		})
	}

	if err := repository.SaveAlumni(*payload); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":      "ALUMNI.SUBMIT.EXCEPTION",
			"message":    "Failed to submit Alumni data!",
			"stacktrace": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Success submit Alumni!",
	})
}

func VerifyAlumni(ctx *fiber.Ctx) error {
	var payload = &model.AlumniVerifyPayload{}
	if err := ctx.BodyParser(payload); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":      "ALUMNI.INVALIDPAYLOAD.EXCEPTION",
			"message":    "Sorry, System can't parse your data! Please Recheck!",
			"stacktrace": err.Error(),
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
			"error":      "ALUMNI.INVALIDPAYLOAD.EXCEPTION",
			"message":    errorMessage,
			"stacktrace": errorMessage,
		})
	}

	id, _ := primitive.ObjectIDFromHex(payload.ID)
	alumniData, err := repository.GetAlumni(bson.M{"_id": id})
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":      "ALUMNI.SUBMIT.EXCEPTION",
			"message":    "Failed to verify Alumni data!",
			"stacktrace": err.Error(),
		})
	}
	if len(alumniData) == 0 {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":      "ALUMNI.SUBMIT.EXCEPTION",
			"message":    "Alumni data not found",
			"stacktrace": "alumni not found",
		})
	}

	if err = repository.VerifyAlumni(*payload); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":      "ALUMNI.SUBMIT.EXCEPTION",
			"message":    "Failed to verify Alumni data!",
			"stacktrace": err.Error(),
		})
	}

	message := fmt.Sprintf(`
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
				<p>Dear Mr/Mrs, %s %s</p>
				<p>Thank you for submit your alumni data. Your data is already verify with Alumni ID: <strong>%s</strong>.</p>
				<p>Thank you,</p>
				<p>Bina Tunas Bangsa School</p>
			</div>
		</body>
		</html>
	`, alumniData[0]["firstname"].(string), alumniData[0]["lastname"].(string), payload.AlumniID)
	var mailPayload pkg.MailPayload = pkg.MailPayload{
		To:      []string{alumniData[0]["email"].(string)},
		Cc:      []string{},
		Subject: "Alumni Verify Notification",
		Message: message,
	}

	err = mailPayload.SendMail()
	if err != nil {
		log.Println(err.Error())
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Success verify Alumni!",
	})
}
