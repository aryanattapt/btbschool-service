package service

import (
	"btb-service/model"
	"btb-service/repository"

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
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":      "CONTACT.INVALIDPAYLOAD.EXCEPTION",
			"message":    "Parameter is required!",
			"stacktrace": err.Error(),
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

	// var mailPayload pkg.MailPayload = pkg.MailPayload{
	// 	To:      []string{payload.Email},
	// 	Cc:      []string{payload.Email},
	// 	Subject: "Contact Submit Notification",
	// 	Message: payload.Message,
	// }

	// err = mailPayload.SendMail()
	// if err != nil {
	// 	return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
	// 		"error":      "CONTACT.SENDMAIL.EXCEPTION",
	// 		"message":    "Failed to send contact data!",
	// 		"stacktrace": err.Error(),
	// 	})
	// }

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Success submit contact!",
	})
}
