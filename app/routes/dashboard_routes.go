package routes

import (
	"log"
	"logistica/app/controllers"
	"logistica/app/models"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

func DashboardRoutes(app *fiber.App, store *session.Store) {
	employee := models.Employee{}
	product := models.Product{}

	app.Get("/dashboard", func(c *fiber.Ctx) error {
		var path string = c.Path()
		var username string = controllers.GetSessionUsername(c, store)

		totalEmployees, err := employee.Count()
		if err != nil {
			log.Println(err)
			return InternalServerError(c, err.Error())
		}

		totalProducts, err := product.Count()
		if err != nil {
			log.Println(err)
			return InternalServerError(c, err.Error())
		}
		// Mengirimkan halaman HTML yang dihasilkan ke browser
		return c.Render("dashboard", fiber.Map{
			"path":           path,
			"user":           username,
			"totalEmployees": totalEmployees,
			"totalProducts":  totalProducts,
		})
	},
	)
}
