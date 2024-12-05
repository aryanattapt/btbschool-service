package main

import (
	"btb-service/controller"
	"btb-service/pkg"
	"btb-service/router"
	"btb-service/service"
	"errors"
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/etag"
	"github.com/gofiber/fiber/v2/middleware/healthcheck"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/robfig/cron"
)

func main() {
	/* Load Env For DEV */
	/* if err := godotenv.Load(); err != nil {
		log.Panic(err.Error())
	} */

	/* Initialize Fiber */
	app := fiber.New(fiber.Config{
		BodyLimit:     128 * 1024 * 1024,
		CaseSensitive: true,
		StrictRouting: false,
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError
			var e *fiber.Error
			if errors.As(err, &e) {
				code = e.Code
			}
			return ctx.Status(code).JSON(fiber.Map{
				"code":    "SERVICE.PANIC.EXCEPTION",
				"message": "Something wrong happened!",
				"error":   err.Error(),
			})
		},
	})

	/* Middleware */
	app.Use(etag.New(etag.ConfigDefault))
	app.Use(logger.New(logger.Config{
		Format: "${ip}:${port} -> ${status} ${method} ${path}\n",
	}))
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowCredentials: false,
		AllowMethods:     "*",
		AllowHeaders:     "*",
		MaxAge:           0,
	}))
	app.Use(recover.New())
	app.Use(healthcheck.New())
	app.Use(helmet.New(helmet.ConfigDefault))
	app.Use(requestid.New(requestid.Config{
		Header: "X-API-Request-ID",
		Generator: func() string {
			return pkg.HashSHA256(pkg.GenerateCurrentTimeStamp())
		},
	}))

	/* Route */
	app.Static("/uploads", os.Getenv("UPLOAD_PATH"))
	app.Static("/assets", os.Getenv("ASSET_PATH"))
	router.CareerRouter(app)
	router.ConfigRouter(app)
	router.AuthRouter(app)
	router.ContactRouter(app)
	router.AlumniRouter(app)
	router.AttachmentsRouter(app)
	router.StudentRegistrationRouter(app)
	router.InstagramRouter(app)
	router.UserRouter(app)
	router.GoogleRouter(app)
	router.SecretKeyRouter(app)
	app.Get("/metrics", monitor.New())
	app.Use(controller.NotFoundRoute)

	/* Cron JOB */
	cron := cron.New()
	cron.AddFunc("@weekly", service.RefreshTokenInstagram)
	cron.Start()

	/* Start Server */
	app.Listen(fmt.Sprintf(":%s", os.Getenv("PORT")))
}
