package controller

import (
	"btb-service/service"

	"github.com/gofiber/fiber/v2"
)

func GetEmailConfig(ctx *fiber.Ctx) error {
	if ctx.Method() == "POST" {
		return service.GetEmailConfig(ctx)
	} else if ctx.Method() == "OPTIONS" {
		return NoContentRoute(ctx)
	} else {
		return MethodNotAllowedRoute(ctx)
	}
}

func UpdateEmailConfig(ctx *fiber.Ctx) error {
	if ctx.Method() == "PUT" {
		return service.UpdateEmailConfig(ctx)
	} else if ctx.Method() == "OPTIONS" {
		return NoContentRoute(ctx)
	} else {
		return MethodNotAllowedRoute(ctx)
	}
}

func GetRecaptchaConfig(ctx *fiber.Ctx) error {
	if ctx.Method() == "POST" {
		return service.GetRecaptchaConfig(ctx)
	} else if ctx.Method() == "OPTIONS" {
		return NoContentRoute(ctx)
	} else {
		return MethodNotAllowedRoute(ctx)
	}
}

func UpdateRecaptchaConfig(ctx *fiber.Ctx) error {
	if ctx.Method() == "PUT" {
		return service.UpdateRecaptchaConfig(ctx)
	} else if ctx.Method() == "OPTIONS" {
		return NoContentRoute(ctx)
	} else {
		return MethodNotAllowedRoute(ctx)
	}
}
