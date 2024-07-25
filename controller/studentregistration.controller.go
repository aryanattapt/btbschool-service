package controller

import (
	"btb-service/service"

	"github.com/gofiber/fiber/v2"
)

func GetStudentRegistrationOutstandingData(ctx *fiber.Ctx) error {
	if ctx.Method() == "POST" {
		return service.GetStudentRegistrationOutstandingData(ctx)
	} else if ctx.Method() == "OPTIONS" {
		return NoContentRoute(ctx)
	} else {
		return MethodNotAllowedRoute(ctx)
	}
}

func SubmitDataStudentRegistration(ctx *fiber.Ctx) error {
	if ctx.Method() == "POST" {
		return service.SubmitDataStudentRegistration(ctx)
	} else if ctx.Method() == "OPTIONS" {
		return NoContentRoute(ctx)
	} else {
		return MethodNotAllowedRoute(ctx)
	}
}
