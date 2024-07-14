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
			"error":      "FETCHCONFIG.INVALIDPAYLOAD.EXCEPTION",
			"message":    "Please define Param!",
			"stacktrace": "Param is not exist",
		})
	}

	var service = repository.ConfigRepositoryModel{ConfigModel: model.ConfigModel{Type: tipe}}
	data, err := service.GetConfig()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":      "FETCHCONFIG.CONFIGQUERY.EXCEPTION",
			"message":    "Failed to get config data!",
			"stacktrace": err.Error(),
		})
	}

	if data == nil || len(data) == 0 {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":      "FETCHCONFIG.CONFIGNOTEXIST.EXCEPTION",
			"message":    "Sorry, we can't find any data. Please try again later!",
			"stacktrace": "Data not found",
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
			"error":      "SAVECONFIG.INVALIDPAYLOAD.EXCEPTION",
			"message":    "Please define Param!",
			"stacktrace": "Param is not exist",
		})
	}

	var payload = &model.ConfigRequestPayload{}
	if err := ctx.BodyParser(payload); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":      "SAVECONFIG.INVALIDPAYLOAD.EXCEPTION",
			"message":    "Sorry, System can't parse your data! Please Recheck!",
			"stacktrace": err.Error(),
		})
	}

	var goValidator = validator.New()
	if err := goValidator.Struct(payload); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":      "SAVECONFIG.INVALIDPAYLOAD.EXCEPTION",
			"message":    "Parameter is required!",
			"stacktrace": err.Error(),
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
			"error":      "SAVECONFIG.UPSERTCONFIG.EXCEPTION",
			"message":    "Failed to save config!",
			"stacktrace": err.Error(),
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
			"error":      "DELETCONFIG.INVALIDPAYLOAD.EXCEPTION",
			"message":    "Please define Param!",
			"stacktrace": "Param is not exist",
		})
	}

	var payload = &model.ConfigRequestPayload{}
	if err := ctx.BodyParser(payload); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":      "DELETECONFIG.INVALIDPAYLOAD.EXCEPTION",
			"message":    "Sorry, System can't parse your data! Please Recheck!",
			"stacktrace": err.Error(),
		})
	}

	var goValidator = validator.New()
	if err := goValidator.Struct(payload); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":      "DELETECONFIG.INVALIDPAYLOAD.EXCEPTION",
			"message":    "Parameter is required!",
			"stacktrace": err.Error(),
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
			"error":      "DELETCONFIG.PURGECONFIG.EXCEPTION",
			"message":    "Failed to delete config!",
			"stacktrace": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Success upsert config",
	})
}
