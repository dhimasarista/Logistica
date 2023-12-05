package routes

import (
	"logistica/app/controllers"
	"logistica/app/models"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

func DashboardRoutes(app *fiber.App, store *session.Store) {
	app.Get("/dashboard", func(c *fiber.Ctx) error {
		var path string = c.Path()
		var username string = controllers.GetSessionUsername(c, store)
		employees := models.Employee{}

		totalEmployees, err := employees.Count()
		if err != nil {
			return InternalServerError(c, err.Error())
		}
		// Mengirimkan halaman HTML yang dihasilkan ke browser
		return c.Render("dashboard", fiber.Map{
			"path":           path,
			"user":           username,
			"totalEmployees": totalEmployees,
		})
	},
	)
}
