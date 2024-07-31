package controller

import (
	"btb-service/service"

	"github.com/gofiber/fiber/v2"
)

func GetDraftStudentRegistrationgData(ctx *fiber.Ctx) error {
	if ctx.Method() == "POST" {
		return service.GetDraftStudentRegistrationData(ctx)
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

func GetStudentRegistrationOutstandingData(ctx *fiber.Ctx) error {
	if ctx.Method() == "GET" {
		return service.GetStudentRegistrationOutstandingData(ctx)
	} else if ctx.Method() == "OPTIONS" {
		return NoContentRoute(ctx)
	} else {
		return MethodNotAllowedRoute(ctx)
	}
}

func ApproveStudentRegistrationOutstandingData(ctx *fiber.Ctx) error {
	if ctx.Method() == "POST" {
		return service.ApproveStudentRegistrationOutstandingData(ctx)
	} else if ctx.Method() == "OPTIONS" {
		return NoContentRoute(ctx)
	} else {
		return MethodNotAllowedRoute(ctx)
	}
}

func GetAllStudentRegistrationAuthData(ctx *fiber.Ctx) error {
	if ctx.Method() == "POST" {
		return service.GetAllStudentRegistrationAuthData(ctx)
	} else if ctx.Method() == "OPTIONS" {
		return NoContentRoute(ctx)
	} else {
		return MethodNotAllowedRoute(ctx)
	}
}

func GetStudentRegistrationDetailData(ctx *fiber.Ctx) error {
	if ctx.Method() == "POST" {
		return service.GetStudentRegistrationDetailData(ctx)
	} else if ctx.Method() == "OPTIONS" {
		return NoContentRoute(ctx)
	} else {
		return MethodNotAllowedRoute(ctx)
	}
}
