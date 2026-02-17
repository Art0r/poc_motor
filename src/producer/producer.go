package producer

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
	"log"
	"time"
)

type Data struct {
	Id        string    `json:"id" validate:"required,uuid4"`
	Email     string    `json:"email" validate:"required,email"`
	Place     string    `json:"placa" validate:"required,min=7,max=8"`
	Timestamp time.Time `json:"timestamp" validate:"required"`
}

var validate = validator.New()

func Start() {
	app := fiber.New()

	app.Post("/", func(c fiber.Ctx) error {

		data := new(Data)

		if err := c.Bind().JSON(&data); err != nil {
			return c.Status(fiber.StatusBadRequest).SendString(err.Error())
		}

		if err := validate.Struct(data); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error":  "Validation failed",
				"fields": err.Error(),
			})
		}

		SendMessage("hello", *data)
		return c.Status(fiber.StatusAccepted).SendString("")
	})

	log.Fatal(app.Listen(":3000"))
}
