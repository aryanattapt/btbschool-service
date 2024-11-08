package controller

import (
	"btb-service/service"

	"github.com/gofiber/fiber/v2"
)

func GetAlumni(ctx *fiber.Ctx) error {
	if ctx.Method() == "POST" {
		return service.GetAlumni(ctx)
	} else if ctx.Method() == "OPTIONS" {
		return NoContentRoute(ctx)
	} else {
		return MethodNotAllowedRoute(ctx)
	}
}

func ValidateAlumni(ctx *fiber.Ctx) error {
	if ctx.Method() == "POST" {
		return service.ValidateAlumni(ctx)
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

func VerifyAlumni(ctx *fiber.Ctx) error {
	if ctx.Method() == "POST" {
		return service.VerifyAlumni(ctx)
	} else if ctx.Method() == "OPTIONS" {
		return NoContentRoute(ctx)
	} else {
		return MethodNotAllowedRoute(ctx)
	}
}

func AlumniAction(ctx *fiber.Ctx) error {
	if ctx.Method() == "DELETE" {
		return service.DeleteAlumni(ctx)
	} else if ctx.Method() == "OPTIONS" {
		return NoContentRoute(ctx)
	} else {
		return MethodNotAllowedRoute(ctx)
	}
}
