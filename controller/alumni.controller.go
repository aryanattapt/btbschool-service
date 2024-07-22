package controller

import (
	"btb-service/service"

	"github.com/gofiber/fiber/v2"
)

func GetAlumni(ctx *fiber.Ctx) error {
	if ctx.Method() == "GET" {
		return service.GetAlumni(ctx)
	} else if ctx.Method() == "OPTIONS" {
		return NoContentRoute(ctx)
	} else {
		return MethodNotAllowedRoute(ctx)
	}
}

func SubmitAlumni(ctx *fiber.Ctx) error {
	if ctx.Method() == "POST" {
		return service.SubmitAlumni(ctx)
	} else if ctx.Method() == "OPTIONS" {
		return NoContentRoute(ctx)
	} else {
		return MethodNotAllowedRoute(ctx)
	}
}
