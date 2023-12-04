package routes

import (
	"fmt"
	"log"
	"logistica/app/controllers"
	"logistica/app/models"
	"strconv"

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
			InternalServerError(c, err.Error())
		}

		return c.Render("employees_page", fiber.Map{
			"path":           path,
			"user":           username,
			"employees":      employees,
			"responseStatus": c.Response().StatusCode(),
		})
	})

	// Memeriksa ketersedian ID
	app.Get("/employee/check/:id", func(c *fiber.Ctx) error {
		var id string = c.Params("id")
		employee := models.Employee{}

		idInteger, err := strconv.Atoi(id) // Konversi string ke integer
		if err != nil {
			log.Println(err)
		}
		employee.GetById(idInteger)

		var isIdExists bool
		if employee.ID.Int64 == int64(idInteger) {
			isIdExists = true
		}

		fmt.Println(employee.ID.Int64)
		fmt.Println(int64(idInteger))
		fmt.Println(employee)

		return c.JSON(fiber.Map{
			"isIdExists":     isIdExists,
			"responseStatus": c.Response().StatusCode(),
		})
	})
	// Mengirim ID baru
	app.Get("/employee/newId", func(c *fiber.Ctx) error {
		employee := models.Employee{}
		err := employee.LastId()
		if err != nil {
			InternalServerError(c, err.Error())
		}

		return c.JSON(fiber.Map{
			"newId":          employee.ID.Int64 + 1,
			"responseStatus": c.Response().StatusCode(),
		})
	})
}
