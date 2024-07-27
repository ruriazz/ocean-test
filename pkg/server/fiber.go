package server

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

type Server interface {
	App() *fiber.App
	Start(port string) error
}

type server struct {
	app *fiber.App
}

func New() (Server, error) {
	app := fiber.New(fiber.Config{
		Prefork:       false,
		CaseSensitive: true,
		StrictRouting: false,
	})

	app.Use(requestid.New(requestid.ConfigDefault))
	app.Use(logger.New(logger.Config{
		Format:     "[${time}][${locals:requestid}] ${status} - ${method} ${path}\n",
		TimeFormat: time.RFC3339Nano,
		TimeZone:   "Asia/Jakarta",
	}))

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("I'm a GET request!")
	})

	return server{
		app: app,
	}, nil
}

func (s server) App() *fiber.App {
	return s.app
}

func (s server) Start(port string) error {
	return s.app.Listen(port)
}
