package service

import (
	"btb-service/model"
	"btb-service/repository"

	"github.com/gofiber/fiber/v2"
)

func GetEmailConfig(ctx *fiber.Ctx) error {
	data, err := repository.GetEmailConfig()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"code":    "FETCHCONFIG.CONFIGQUERY.EXCEPTION",
			"message": "Failed to get email config dataa!",
			"error":   err.Error(),
		})
	}

	if len(data) == 0 {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"code":    "FETCHCONFIG.CONFIGNOTEXIST.EXCEPTION",
			"message": "Sorry, we can't find any email config data. Please try again later!",
			"error":   "Data not found",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Success get email config",
		"data":    data,
	})
}

func UpdateEmailConfig(ctx *fiber.Ctx) error {
	var payload = &model.SecretKeyPayload{}
	if err := ctx.BodyParser(payload); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"code":    "SAVECONFIG.INVALIDPAYLOAD.EXCEPTION",
			"message": "Sorry, System can't parse your data! Please Recheck!",
			"error":   err.Error(),
		})
	}

	if err := repository.UpdateEmailConfig(*payload); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"code":    "SAVECONFIG.UPSERTCONFIG.EXCEPTION",
			"message": "Failed to save email config!",
			"error":   err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Success save email config",
	})
}

func GetRecaptchaConfig(ctx *fiber.Ctx) error {
	data, err := repository.GetRecaptchaConfig()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"code":    "FETCHCONFIG.CONFIGQUERY.EXCEPTION",
			"message": "Failed to get recaptcha config data!",
			"error":   err.Error(),
		})
	}

	if len(data) == 0 {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"code":    "FETCHCONFIG.CONFIGNOTEXIST.EXCEPTION",
			"message": "Sorry, we can't find any recaptcha config data. Please try again later!",
			"error":   "Data not found",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Success get recaptcha config",
		"data":    data,
	})
}

func UpdateRecaptchaConfig(ctx *fiber.Ctx) error {
	var payload = &model.SecretKeyPayload{}
	if err := ctx.BodyParser(payload); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"code":    "SAVECONFIG.INVALIDPAYLOAD.EXCEPTION",
			"message": "Sorry, System can't parse your data! Please Recheck!",
			"error":   err.Error(),
		})
	}

	if err := repository.UpdateEmailConfig(*payload); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"code":    "SAVECONFIG.UPSERTCONFIG.EXCEPTION",
			"message": "Failed to save recaptcha config!",
			"error":   err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Success save recaptcha config",
	})
}
