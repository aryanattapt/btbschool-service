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
			"message": "Please define type!",
			"error":   "INVALID_PAYLOAD",
		})
	}

	var service = repository.ConfigRepositoryModel{ConfigModel: model.ConfigModel{Type: tipe}}
	data, err := service.GetConfig()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
			"error":   "QUERY_EXCEPTION",
		})
	}

	if data == nil || len(data) == 0 {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Sorry, we can't find any data. Please try again later!",
			"error":   "DATA_NOT_FOUND",
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
			"message": "Please define type!",
			"error":   "INVALID_PAYLOAD",
		})
	}

	var payload = &model.ConfigRequestPayload{}
	if err := ctx.BodyParser(payload); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Sorry, System can't parse your data! Please Recheck!",
			"error":   "INVALID_PAYLOAD",
		})
	}

	var goValidator = validator.New()
	if err := goValidator.Struct(payload); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
			"error":   "INVALID_PAYLOAD",
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
			"message": err.Error(),
			"error":   "QUERY_EXCEPTION",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Success upsert config",
	})
}

func DeleteConfig(ctx *fiber.Ctx) error {
	var tipe string = ctx.Params("type")
	if pkg.IsEmptyString(tipe) {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Please define type!",
			"error":   "INVALID_PAYLOAD",
		})
	}

	var payload = &model.ConfigRequestPayload{}
	if err := ctx.BodyParser(payload); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Sorry, System can't parse your data! Please Recheck!",
			"error":   "INVALID_PAYLOAD",
		})
	}

	var goValidator = validator.New()
	if err := goValidator.Struct(payload); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
			"error":   "INVALID_PAYLOAD",
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
			"message": err.Error(),
			"error":   "QUERY_EXCEPTION",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Success upsert config",
	})
}
