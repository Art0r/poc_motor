package producer

import (
	"log"

	"github.com/gofiber/fiber/v3"
)

func Start() {
	app := fiber.New()

	app.Get("/", func(c fiber.Ctx) error {
		SendMessage("hello", "Ola Mundo")
		return c.SendString("Producer")
	})

	log.Fatal(app.Listen(":3000"))
}
