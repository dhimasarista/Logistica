package routes

import (
	"logistica/app/controllers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

func ReportsRoutes(app *fiber.App, store *session.Store) {

	app.Get("/reports", func(c *fiber.Ctx) error {
		// Mendapatkan path dan username dari sesi
		var path string = c.Path()
		var username string = controllers.GetSessionUsername(c, store)
		return c.Render("reports_page", fiber.Map{
			"path": path,
			"user": username,
		})
	})
}
