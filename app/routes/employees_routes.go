package routes

import (
	"logistica/app/controllers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

func EmployeesRoutes(app *fiber.App, store *session.Store) {
	app.Get("/employees", func(c *fiber.Ctx) error {
		var path string = c.Path()
		var username string = controllers.GetSessionUsername(c, store)

		var data = models.Employee{}
		fmt.Println(data.GetById(1))

		return c.Render("employees_page", fiber.Map{
			"path": path,
			"user": username,
		})
	})
}
