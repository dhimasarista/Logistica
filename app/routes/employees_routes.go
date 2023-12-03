package routes

import (
	"log"
	"logistica/app/controllers"
	"logistica/app/models"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

func EmployeesRoutes(app *fiber.App, store *session.Store) {
	app.Get("/employees", func(c *fiber.Ctx) error {
		var path string = c.Path()
		var username string = controllers.GetSessionUsername(c, store)

		employee := models.Employee{}

		employees, err := employee.FindAll()
		if err != nil {
			log.Println(err)
			InternalServerError(c, err.Error())
		}

		return c.Render("employees_page", fiber.Map{
			"path":           path,
			"user":           username,
			"employees":      employees,
			"responseStatus": c.Response().StatusCode(),
		})
	})
}
