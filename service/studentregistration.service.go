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
			"code":    "STUDENTREG.INVALIDPAYLOAD.EXCEPTION",
			"message": "Sorry, System can't parse your data! Please Recheck!",
			"error":   err.Error(),
		})
	}

	data, err := repository.GetDraftStudentRegistrationData(*payload)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"code":    "STUDENTREG.QUERY.EXCEPTION",
			"message": "Failed to get Student Registration Outstanding data!",
			"error":   err.Error(),
		})
	}

	if len(data) == 0 {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"code":    "STUDENTREG.NOTEXIST.EXCEPTION",
			"message": "Sorry, Student Registration Outstanding Data isn't exist!",
			"error":   "Student Registration Outstanding data is not exist",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Success get Student Registration Outstanding data!",
		"data":    data,
	})
}

func parseAndValidatePage(ctx *fiber.Ctx, payload interface{}, pageNumber int) map[string]interface{} {
	var goValidator = validator.New()
	errorResultValidator := make(map[string]interface{})

	if err := ctx.BodyParser(payload); err != nil {
		ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"code":    "STUDENTREG.INVALIDPAYLOAD.EXCEPTION",
			"message": "Sorry, System can't parse your data! Please Recheck!",
			"error":   err.Error(),
		})
		return nil
	}

	if err := goValidator.Struct(payload); err != nil {
		errorResult := pkg.ValidateForOnlineRegistration(err, pageNumber)
		for key, value := range errorResult {
			errorResultValidator[key] = value
		}
	}
	return errorResultValidator
}

func ValidateDataStudentRegistration(ctx *fiber.Ctx) error {
	combinedErrors := make(map[string]interface{})

	/* Page 1 */
	payloadPage1 := &model.StudentRegistrationInsertPayloadPage1{}
	page1Errors := parseAndValidatePage(ctx, payloadPage1, 1)
	if len(page1Errors) > 0 {
		for key, value := range page1Errors {
			combinedErrors[key] = value
		}
	}

	/* Page 3 */
	payloadPage3 := &model.StudentRegistrationInsertPayloadPage3{}
	page3Errors := parseAndValidatePage(ctx, payloadPage3, 3)
	if len(page3Errors) > 0 {
		for key, value := range page3Errors {
			combinedErrors[key] = value
		}
	}

	/* Page 4 */
	payloadPage4 := &model.StudentRegistrationInsertPayloadPage4{}
	page4Errors := parseAndValidatePage(ctx, payloadPage4, 4)
	if len(page4Errors) > 0 {
		for key, value := range page4Errors {
			combinedErrors[key] = value
		}
	}

	if len(combinedErrors) > 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"code":    "STUDENTREG.INVALIDPAYLOAD.EXCEPTION",
			"message": "Validation errors occurred!",
			"error":   combinedErrors,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "All data is valid!",
	})
}

func SubmitDataStudentRegistration(ctx *fiber.Ctx) error {
	var payload = &model.StudentRegistrationInsertPayload{}
	if err := ctx.BodyParser(payload); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"code":    "STUDENTREG.INVALIDPAYLOAD.EXCEPTION",
			"message": "Sorry, System can't parse your data! Please Recheck!",
			"error":   err.Error(),
		})
	}

	if payload.Status == "send" {
		combinedErrors := make(map[string]interface{})

		/* Page 1 */
		payloadPage1 := &model.StudentRegistrationInsertPayloadPage1{}
		page1Errors := parseAndValidatePage(ctx, payloadPage1, 1)
		if len(page1Errors) > 0 {
			for key, value := range page1Errors {
				combinedErrors[key] = value
			}
		}

		/* Page 3 */
		payloadPage3 := &model.StudentRegistrationInsertPayloadPage3{}
		page3Errors := parseAndValidatePage(ctx, payloadPage3, 3)
		if len(page3Errors) > 0 {
			for key, value := range page3Errors {
				combinedErrors[key] = value
			}
		}

		/* Page 4 */
		payloadPage4 := &model.StudentRegistrationInsertPayloadPage4{}
		page4Errors := parseAndValidatePage(ctx, payloadPage4, 4)
		if len(page4Errors) > 0 {
			for key, value := range page4Errors {
				combinedErrors[key] = value
			}
		}

		if len(combinedErrors) > 0 {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"code":    "STUDENTREG.INVALIDPAYLOAD.EXCEPTION",
				"message": "Validation errors occurred!",
				"error":   combinedErrors,
			})
		}
	}

	registrationCode, err := repository.SubmitDataStudentRegistration(*payload)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"code":    "STUDENTREG.SUBMIT.EXCEPTION",
			"message": "Failed to submit student data!",
			"error":   err.Error(),
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
		To:      []string{payload.MainEmail},
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
			"code":    "STUDENTREG.QUERY.EXCEPTION",
			"message": "Failed to get outstanding data!",
			"error":   err.Error(),
		})
	}

	if len(data) == 0 {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"code":    "STUDENTREG.NOTEXIST.EXCEPTION",
			"message": "Sorry, Student Registration Outstanding Data isn't exist!",
			"error":   "Student Registration Outstanding data is not exist",
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
			"code":    "STUDENTREG.INVALIDPAYLOAD.EXCEPTION",
			"message": "Sorry, System can't parse your data! Please Recheck!",
			"error":   err.Error(),
		})
	}

	err := repository.ApproveStudentRegistrationOutstandingData(ctx.Locals("userauth"), *payload)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"code":    "STUDENTREG.SUBMIT.EXCEPTION",
			"message": "Failed to submit data!",
			"error":   err.Error(),
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
			"code":    "STUDENTREG.QUERY.EXCEPTION",
			"message": "Failed to get all student data!",
			"error":   err.Error(),
		})
	}

	if len(data) == 0 {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"code":    "STUDENTREG.NOTEXIST.EXCEPTION",
			"message": "Sorry, Student Registration Data isn't exist!",
			"error":   "Student Registration data is not exist",
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
			"code":    "STUDENTREG.INVALIDPAYLOAD.EXCEPTION",
			"message": "Sorry, System can't parse your data! Please Recheck!",
			"error":   err.Error(),
		})
	}

	data, err := repository.GetStudentRegistrationDetailData(*payload)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"code":    "STUDENTREG.QUERY.EXCEPTION",
			"message": "Failed to get student data!",
			"error":   err.Error(),
		})
	}

	if len(data) == 0 {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"code":    "STUDENTREG.NOTEXIST.EXCEPTION",
			"message": "Sorry, Student Registration Data isn't exist!",
			"error":   "Student Registration data is not exist",
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Success get student data!",
		"data":    data,
	})
}
