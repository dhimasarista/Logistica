package routes

import (
	"logistica/app/controllers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

func DashboardRoutes(app *fiber.App, store *session.Store) {
	app.Get("/dashboard", func(c *fiber.Ctx) error {
		var path string = c.Path()
		var username string = controllers.GetSessionUsername(c, store)

		// Mengirimkan halaman HTML yang dihasilkan ke browser
		return c.Render("dashboard", fiber.Map{
			"path": path,
			"user": username,
		})
	},
	)
}
