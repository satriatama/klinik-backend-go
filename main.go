package main

import (
	pasiencontrollers "github.com/satriatama/klinik-backend-go/controllers/pasienControllers"

	"github.com/gofiber/fiber/v2"
	"github.com/satriatama/klinik-backend-go/models"
)

func main() {
	models.ConnectDatabase()
	app := fiber.New()


	app.Get("/api/v1/pasien", pasiencontrollers.Index)

	app.Listen(":8000")
}