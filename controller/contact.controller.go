package controller

import (
	"btb-service/service"

	"github.com/gofiber/fiber/v2"
)

func GetContacts(ctx *fiber.Ctx) error {
	if ctx.Method() == "GET" {
		return service.GetContacts(ctx)
	} else if ctx.Method() == "OPTIONS" {
		return NoContentRoute(ctx)
	} else {
		return MethodNotAllowedRoute(ctx)
	}
}

func SaveContacts(ctx *fiber.Ctx) error {
	if ctx.Method() == "POST" {
		return service.SubmitContact(ctx)
	} else if ctx.Method() == "OPTIONS" {
		return NoContentRoute(ctx)
	} else {
		return MethodNotAllowedRoute(ctx)
	}
}
