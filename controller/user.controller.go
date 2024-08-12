package controller

import (
	"btb-service/service"

	"github.com/gofiber/fiber/v2"
)

func GetAllUser(ctx *fiber.Ctx) error {
	if ctx.Method() == "POST" {
		return service.GetAllUser(ctx)
	} else if ctx.Method() == "OPTIONS" {
		return NoContentRoute(ctx)
	} else {
		return MethodNotAllowedRoute(ctx)
	}
}

func UpsertUser(ctx *fiber.Ctx) error {
	if ctx.Method() == "POST" {
		return service.InsertUser(ctx)
	} else if ctx.Method() == "PUT" || ctx.Method() == "PATCH" {
		return service.UpdateUser(ctx)
	} else if ctx.Method() == "OPTIONS" {
		return NoContentRoute(ctx)
	} else {
		return MethodNotAllowedRoute(ctx)
	}
}
