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

func GetContacts(ctx *fiber.Ctx) error {
	data, err := repository.GetContacts()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":      "CONTACT.CONTACTQUERY.EXCEPTION",
			"message":    "Failed to get contact data!",
			"stacktrace": err.Error(),
		})
	}

	if len(data) == 0 {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":      "CONTACT.NOTEXIST.EXCEPTION",
			"message":    "Sorry, Contact Data isn't exist!",
			"stacktrace": "Contact data is not exist",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Success get contact!",
		"data":    data,
	})
}

func SubmitContact(ctx *fiber.Ctx) error {
	var payload = &model.ContactInsertPayload{}
	if err := ctx.BodyParser(payload); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":      "CONTACT.INVALIDPAYLOAD.EXCEPTION",
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
			"error":      "CONTACT.INVALIDPAYLOAD.EXCEPTION",
			"message":    errorMessage,
			"stacktrace": errorMessage,
		})
	}

	err := repository.SaveContacts(*payload)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":      "CONTACT.SUBMIT.EXCEPTION",
			"message":    "Failed to submit contact data!",
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
				<p>Dear Mr/Mrs Bina Tunas Bangsa School,</p>
				<p>You just received a Message From %s %s. The Message are below.</p>
				<strong>%s</strong>
				<p>Thank you</p>
			</div>
		</body>
		</html>
	`, payload.FirstName, payload.LastName, payload.Message)

	var mailPayload pkg.MailPayload = pkg.MailPayload{
		To:      []string{"aryanatta@gmail.com"},
		Cc:      []string{payload.Email},
		Subject: "Contact Submit Notification",
		Message: message,
	}

	err = mailPayload.SendMail()
	if err != nil {
		log.Println(err.Error())
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Success submit contact!",
	})
}
