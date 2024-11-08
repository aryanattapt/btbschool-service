package service

import (
	"btb-service/model"
	"btb-service/pkg"
	"btb-service/repository"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func GetConfigs(ctx *fiber.Ctx) error {
	var tipe string = ctx.Params("type")
	if pkg.IsEmptyString(tipe) {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"code":    "FETCHCONFIG.INVALIDPAYLOAD.EXCEPTION",
			"message": "Please define Param!",
			"error":   "Param is not exist",
		})
	}

	var payload = &map[string]interface{}{}
	if err := ctx.BodyParser(payload); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"code":    "FETCHCONFIG.INVALIDPAYLOAD.EXCEPTION",
			"message": "Sorry, System can't parse your data! Please Recheck!",
			"error":   err.Error(),
		})
	}

	var service = repository.ConfigRepositoryModel{ConfigModel: model.ConfigModel{Type: tipe}}
	data, err := service.GetConfig(*payload)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"code":    "FETCHCONFIG.CONFIGQUERY.EXCEPTION",
			"message": "Failed to get config data!",
			"error":   err.Error(),
		})
	}

	if len(data) == 0 {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"code":    "FETCHCONFIG.CONFIGNOTEXIST.EXCEPTION",
			"message": "Sorry, we can't find any data. Please try again later!",
			"error":   "Data not found",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Success get config",
		"data":    data,
	})
}

func UpsertConfig(ctx *fiber.Ctx) error {
	var tipe string = ctx.Params("type")
	if pkg.IsEmptyString(tipe) {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"code":    "SAVECONFIG.INVALIDPAYLOAD.EXCEPTION",
			"message": "Please define Param!",
			"error":   "Param is not exist",
		})
	}

	var payload = &model.ConfigRequestPayload{}
	if err := ctx.BodyParser(payload); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"code":    "SAVECONFIG.INVALIDPAYLOAD.EXCEPTION",
			"message": "Sorry, System can't parse your data! Please Recheck!",
			"error":   err.Error(),
		})
	}

	var goValidator = validator.New()
	if err := goValidator.Struct(payload); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"code":    "SAVECONFIG.INVALIDPAYLOAD.EXCEPTION",
			"message": "Parameter is required!",
			"error":   err.Error(),
		})
	}

	m := make([]map[string]interface{}, 0)
	for _, v := range payload.Payload {
		val, ok := v.(map[string]interface{})
		if ok {
			m = append(m, val)
		}
	}

	var service = repository.ConfigRepositoryModel{ConfigModel: model.ConfigModel{Type: tipe, Payload: m}}
	if err := service.UpsertConfig(); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"code":    "SAVECONFIG.UPSERTCONFIG.EXCEPTION",
			"message": "Failed to save config!",
			"error":   err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Success save config",
	})
}

func DeleteConfig(ctx *fiber.Ctx) error {
	var tipe string = ctx.Params("type")
	if pkg.IsEmptyString(tipe) {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"code":    "DELETCONFIG.INVALIDPAYLOAD.EXCEPTION",
			"message": "Please define Param!",
			"error":   "Param is not exist",
		})
	}

	var payload = &model.ConfigRequestPayload{}
	if err := ctx.BodyParser(payload); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"code":    "DELETECONFIG.INVALIDPAYLOAD.EXCEPTION",
			"message": "Sorry, System can't parse your data! Please Recheck!",
			"error":   err.Error(),
		})
	}

	var goValidator = validator.New()
	if err := goValidator.Struct(payload); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"code":    "DELETECONFIG.INVALIDPAYLOAD.EXCEPTION",
			"message": "Parameter is required!",
			"error":   err.Error(),
		})
	}

	m := make([]map[string]interface{}, 0)
	for _, v := range payload.Payload {
		val, ok := v.(map[string]interface{})
		if ok {
			m = append(m, val)
		}
	}

	var service = repository.ConfigRepositoryModel{ConfigModel: model.ConfigModel{Type: tipe, Payload: m}}
	if err := service.DeleteConfig(); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"code":    "DELETCONFIG.PURGECONFIG.EXCEPTION",
			"message": "Failed to delete config!",
			"error":   err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Success delete config",
	})
}
