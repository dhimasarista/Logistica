package routes

import (
	"logistica/app/controllers"

	"github.com/gofiber/fiber/v2"
)

func IndexRoutes(app *fiber.App) {
	app.Get("/home", controllers.IndexRender)
	app.Get("/", func(c *fiber.Ctx) error {

		return c.Redirect("/home")
	})
}
