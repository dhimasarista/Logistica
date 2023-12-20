package routes

import (
	"logistica/app/controllers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

func OrdersRoutes(app *fiber.App, store *session.Store) {
	app.Get("/orders", func(c *fiber.Ctx) error {
		var path string = c.Path()
		var username string = controllers.GetSessionUsername(c, store)

		return c.Render("orders_page", fiber.Map{
			"path": path,
			"user": username,
		})
	})
}
