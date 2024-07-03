package controller

import (
	"btb-service/service"

	"github.com/gofiber/fiber/v2"
)

func ConfigController(ctx *fiber.Ctx) error {
	if ctx.Method() == "GET" {
		return service.GetConfigs(ctx)
	} else if ctx.Method() == "POST" || ctx.Method() == "PUT" || ctx.Method() == "PATCH" {
		return service.UpsertConfig(ctx)
	} else if ctx.Method() == "DELETE" {
		return service.DeleteConfig(ctx)
	} else if ctx.Method() == "OPTIONS" {
		return NoContentRoute(ctx)
	} else {
		return MethodNotAllowedRoute(ctx)
	}
}
