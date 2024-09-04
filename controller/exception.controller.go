package controller

import "github.com/gofiber/fiber/v2"

func MethodNotAllowedRoute(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusMethodNotAllowed).JSON(fiber.Map{
		"code":    "ROUTER.METHODNOTALLOWED.EXCEPTION",
		"message": "Sorry, method is not allowed in this URL!",
	})
}

func NotFoundRoute(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
		"code":    "ROUTER.ROUTENOTFOUND.EXCEPTION",
		"message": "Sorry, destination not found",
	})
}

func NoContentRoute(ctx *fiber.Ctx) error {
	ctx.Status(fiber.StatusNoContent)
	return nil
}
