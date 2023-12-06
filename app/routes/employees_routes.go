package routes

import (
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
		position := models.Position{}

		employees, err := employee.FindAll()
		if err != nil {
			InternalServerError(c, err.Error())
		}

		positions, err := position.FindAll()
		if err != nil {
			InternalServerError(c, err.Error())
		}

		return c.Render("employees_page", fiber.Map{
			"path":           path,
			"user":           username,
			"employees":      employees,
			"positions":      positions,
			"responseStatus": c.Response().StatusCode(),
		})
	})

	// Memeriksa ketersedian ID
	app.Get("/employee/check/:id", func(c *fiber.Ctx) error {
		var id string = c.Params("id")
		employee := &models.Employee{}

		idInteger, err := strconv.Atoi(id) // Konversi string ke integer
		if err != nil {
			log.Println(err)
		}
		employee.GetById(int64(idInteger)) // Mengambil ID

		var isIdExists bool = false
		if employee.ID.Int64 == int64(idInteger) {
			isIdExists = true
		}

		return c.JSON(fiber.Map{
			"isIdExists":     isIdExists,
			"responseStatus": c.Response().StatusCode(),
		})
	})
	// Mengirim ID baru
	app.Get("/employee/newId", func(c *fiber.Ctx) error {
		employee := models.Employee{}
		lastId, err := employee.LastId()
		if err != nil {
			InternalServerError(c, err.Error())
		}

		return c.JSON(fiber.Map{
			"newId":          lastId + 1,
			"responseStatus": c.Response().StatusCode(),
		})
	})

	app.Post("/employee/new", func(c *fiber.Ctx) error {
		var formData map[string]interface{}
		var employee = models.Employee{}

		err := c.BodyParser(&formData)
		if err != nil {
			log.Println(err)
			return c.JSON(fiber.Map{
				"error":  err.Error(),
				"status": 500,
			})
		}

		if formData["id"] == "" || formData["name"] == "" || formData["numberPhone"] == "" || formData["position"] == "" {
			return c.JSON(fiber.Map{
				"error":  "Form is Empty",
				"status": 400,
			})
		}

		idToInt, _ := strconv.Atoi(formData["id"].(string))
		positionToInt, _ := strconv.Atoi(formData["position"].(string))
		var isUser int = 0
		if formData["isUser"].(bool) {
			isUser = 1
		}
		newEmpResult, err := employee.NewEmployee(
			idToInt,
			formData["name"].(string),
			formData["address"].(string),
			formData["numberPhone"].(string),
			positionToInt,
			isUser,
		)

		if err != nil {
			log.Println(err)
			return c.JSON(fiber.Map{
				"error":  err.Error(),
				"status": 500,
			})
		}

		return c.JSON(fiber.Map{
			"error":  nil,
			"status": c.Response().StatusCode(),
			"data":   formData,
			"result": newEmpResult,
		})
	})

	app.Post("/employee/position/new", func(c *fiber.Ctx) error {
		var position = models.Position{}
		var formData map[string]any

		err := c.BodyParser(&formData)
		if err != nil {
			log.Println(err)
			return c.JSON(fiber.Map{
				"error":  err.Error(),
				"status": 500,
			})
		}

		if formData["name"].(string) == "" {
			return c.JSON(fiber.Map{
				"error":  "Position Filed is Empty!",
				"status": 400,
			})
		}

		lastId, err := position.LastId()
		if err != nil {
			log.Println(err)
			return err
		}
		result, err := position.NewPosition(lastId+1, formData["name"].(string))
		if err != nil {
			log.Println(err)
			return c.JSON(fiber.Map{
				"error":  err.Error(),
				"status": 500,
			})
		}

		return c.JSON(fiber.Map{
			"error":  nil,
			"status": c.Response().StatusCode(),
			"data":   formData,
			"result": result,
		})

	})
}
