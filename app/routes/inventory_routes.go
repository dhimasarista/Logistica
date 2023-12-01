package routes

import (
	"logistica/app/controllers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

func InventoryRoutes(app *fiber.App, store *session.Store) {
	app.Get("/inventory", func(c *fiber.Ctx) error {
		var path string = c.Path()
		var username string = controllers.GetSessionUsername(c, store)

		return c.Render("inventory_page", fiber.Map{
			"path": path,
			"user": username,
		})
	},
	)
}
