package main

import (
	"github.com/gofiber/fiber"
)

func api(c *fiber.Ctx) {
	c.Send("KATHIR!")
}

func setupRoutes(app *fiber.App) {
	app.Get("/", api)

}

func main() {
	app := fiber.New()

	setupRoutes(app)
	app.Listen(8888)
}
