package controller

import "github.com/gofiber/fiber/v2"

func MethodNotAllowedRoute(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusMethodNotAllowed).JSON(fiber.Map{
		"message": "Sorry, method is not allowed in this URL!",
		"error":   "METHOD_NOT_ALLOWED",
	})
}

func NotFoundRoute(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
		"message": "Sorry, destination not found",
		"error":   "ROUTE_NOT_FOUND",
	})
}

func NoContentRoute(ctx *fiber.Ctx) error {
	ctx.Status(fiber.StatusNoContent)
	return nil
}
