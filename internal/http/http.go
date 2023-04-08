package http

import (
	_ "video/docs"

	"video/internal/http/controllers"

	fiberSwagger "github.com/swaggo/fiber-swagger"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/monitor"
)

// HealthCheck godoc
// @Summary Show the status of server.
// @Description get the status of server.
// @Tags root
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router / [get]
func healthCheck(c *fiber.Ctx) error {
	res := map[string]interface{}{
		"data": "Server is up and running",
	}

	if err := c.JSON(res); err != nil {
		return err
	}

	return nil
}

func Http() *fiber.App {
	app := fiber.New(fiber.Config{
		StreamRequestBody: true,
		BodyLimit:         100 * 1024 * 1024 * 1024,
	})

	app.Use(cors.New(cors.Config{
		Next:             nil,
		AllowOrigins:     "*",
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH",
		AllowHeaders:     "",
		AllowCredentials: false,
		ExposeHeaders:    "",
		MaxAge:           0,
	}))

	app.Use(compress.New(compress.Config{
		Level: compress.LevelBestSpeed,
	}))

	app.Static("data", "./data")

	// app.Get("/checkUser/:username", CheckUserName)
	app.Get("/upload", healthCheck)
	app.Get("/upload/swagger/*", fiberSwagger.FiberWrapHandler())

	app.Get("/upload/monitor", monitor.New(monitor.Config{
		Title: "Upload Monitor",
	}))

	controllers.Uploader(app.Group("/files"))

	return app
}
