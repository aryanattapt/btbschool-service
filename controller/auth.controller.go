package controller

import (
	"btb-service/service"

	"github.com/gofiber/fiber/v2"
)

func SignUp(ctx *fiber.Ctx) error {
	if ctx.Method() == "POST" {
		return service.SignUp(ctx)
	} else if ctx.Method() == "OPTIONS" {
		return NoContentRoute(ctx)
	} else {
		return MethodNotAllowedRoute(ctx)
	}
}

func SignIn(ctx *fiber.Ctx) error {
	if ctx.Method() == "POST" {
		return service.SignIn(ctx)
	} else if ctx.Method() == "OPTIONS" {
		return NoContentRoute(ctx)
	} else {
		return MethodNotAllowedRoute(ctx)
	}
}

func Validate(ctx *fiber.Ctx) error {
	if ctx.Method() == "POST" {
		return service.Validate(ctx)
	} else if ctx.Method() == "OPTIONS" {
		return NoContentRoute(ctx)
	} else {
		return MethodNotAllowedRoute(ctx)
	}
}

func GetOrUpdateUser(ctx *fiber.Ctx) error {
	if ctx.Method() == "POST" {
		return service.GetAllUser(ctx)
	} else if ctx.Method() == "PUT" || ctx.Method() == "PATCH" {
		return service.UpdateUser(ctx)
	} else if ctx.Method() == "OPTIONS" {
		return NoContentRoute(ctx)
	} else {
		return MethodNotAllowedRoute(ctx)
	}
}
