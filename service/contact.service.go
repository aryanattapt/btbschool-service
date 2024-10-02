package service

import (
	"btb-service/model"
	"btb-service/pkg"
	"btb-service/repository"
	"fmt"
	"html/template"
	"log"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func GetContacts(ctx *fiber.Ctx) error {
	data, err := repository.GetContacts()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"code":    "CONTACT.CONTACTQUERY.EXCEPTION",
			"message": "Failed to get contact data!",
			"error":   err.Error(),
		})
	}

	if len(data) == 0 {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"code":    "CONTACT.NOTEXIST.EXCEPTION",
			"message": "Sorry, Contact Data isn't exist!",
			"error":   "Contact data is not exist",
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
			"code":    "CONTACT.INVALIDPAYLOAD.EXCEPTION",
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
			"code":    "CONTACT.INVALIDPAYLOAD.EXCEPTION",
			"message": errorMessage,
			"error":   errorMessage,
		})
	}

	err := repository.SaveContacts(*payload)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"code":    "CONTACT.SUBMIT.EXCEPTION",
			"message": "Failed to submit contact data!",
			"error":   err.Error(),
		})
	}

	var service = repository.ConfigRepositoryModel{ConfigModel: model.ConfigModel{Type: "general"}}
	data, err := service.GetConfig(map[string]interface{}{
		"type": "admincms.contact.mailcontent",
	})
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"code":    "FETCHCONFIG.CONFIGQUERY.EXCEPTION",
			"message": "Failed to get config data!",
			"error":   err.Error(),
		})
	}

	if len(data) == 0 {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"code":    "FETCHCONFIG.CONFIGQUERY.EXCEPTION",
			"message": "Failed to get mail content setting!",
		})
	}

	mailContent, ok := data[0]["content"].(string)
	if !ok {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"code":    "FETCHCONFIG.CONFIGQUERY.EXCEPTION",
			"message": "Failed to parse mail content setting!",
		})
	}
	mailContent = fmt.Sprintf(mailContent, payload.FirstName, payload.LastName, payload.Message)

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
			%s
		</body>
		</html>
	`, template.HTML(mailContent))

	emailString := data[0]["emailrecepient"].(string)
	emails := strings.Split(emailString, ";")
	for i := range emails {
		emails[i] = strings.TrimSpace(emails[i])
	}

	var mailPayload pkg.MailPayload = pkg.MailPayload{
		To:      []string{payload.Email},
		Cc:      emails,
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
