package web

import (
	scheduler "cth.release/go-rds/cron"
	"cth.release/go-rds/web/channel"
	"cth.release/go-rds/web/storage"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/robfig/cron/v3"
)

type ServerConfig struct {
	App  *fiber.App
	Cron *cron.Cron
}

func InitServer() *ServerConfig {
	app := fiber.New()

	c := scheduler.InitCron()

	server := &ServerConfig{
		App:  app,
		Cron: c,
	}

	server.SetupRoutes(app)
	return server
}

func (s *ServerConfig) SetupRoutes(app *fiber.App) *fiber.App {
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,PUT,DELETE",
		AllowHeaders: "Origin, Content-Type",
	}))

	api := app.Group("/api", EmptyMiddleware)
	api.Get("/health", HealthHandler)

	channel.SetupRoutes(api)
	storage.SetupRoutes(api)

	return app
}

func EmptyMiddleware(c *fiber.Ctx) error {
	return c.Next()
}

func HealthHandler(c *fiber.Ctx) error {
	return c.SendString("200")
}
