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
	/*
		Inisialisasi variabel employee sebagai pointer ke struct models.Employee.
		Dengan menggunakan pointer, kita dapat memanfaatkan referensi ke data Employee yang sudah ada
		dan menghindari pembuatan salinan nilai data, yang dapat menghemat memori.
	*/
	var employee *models.Employee = &models.Employee{}
	position := models.Position{}

	// Halaman-halaman yang dirender
	app.Get("/employees", func(c *fiber.Ctx) error {
		var path string = c.Path()
		var username string = controllers.GetSessionUsername(c, store)

		employees, err := employee.FindAll()
		if err != nil {
			log.Println(err)
			InternalServerError(c, err.Error())
		}

		positions, err := position.FindAll()
		if err != nil {
			log.Println(err)
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

	// Menghapus employee berdasarkan id
	app.Delete("/employee/delete/:id", func(c *fiber.Ctx) error {
		var id string = c.Params("id")

		idInteger, err := strconv.Atoi(id)
		if err != nil {
			log.Println(err)
		}

		err = employee.DeleteEmployee(idInteger)
		if err != nil {
			log.Println(err)
			return c.JSON(fiber.Map{
				"error":          err.Error(),
				"responseStatus": 500,
			})
		}

		return c.JSON(fiber.Map{
			"responseStatus": c.Response().StatusCode(),
			"message":        "Employee Deleted!",
		})

	})

	// Mengirim ID baru
	app.Get("/employee/newId", func(c *fiber.Ctx) error {
		lastId, err := employee.LastId()
		if lastId <= 100020 {
			lastId = 100020
		}

		if err != nil {
			log.Println(err)
			InternalServerError(c, err.Error())
		}

		return c.JSON(fiber.Map{
			"newId":          lastId + 1,
			"responseStatus": c.Response().StatusCode(),
		})
	})

	// Membuat data employee baru
	app.Post("/employee/new", func(c *fiber.Ctx) error {
		var formData map[string]interface{} // variabel untuk menyimpan data yang diterima dari client-side
		err := c.BodyParser(&formData)
		if err != nil {
			log.Println(err)
			return c.JSON(fiber.Map{
				"error":  err.Error(),
				"status": fiber.StatusInternalServerError,
			})
		}

		if formData["id"] == "" || formData["name"] == "" || formData["numberPhone"] == "" || formData["position"] == "" {
			return c.JSON(fiber.Map{
				"error":  "Form is Empty",
				"status": fiber.StatusBadRequest,
			})
		}

		idToInt, err := strconv.Atoi(formData["id"].(string))
		if err != nil {
			panic(err)
		}

		positionToInt, err := strconv.Atoi(formData["position"].(string))
		if err != nil {
			panic(err)
		}

		newEmpResult, err := employee.NewEmployee(
			idToInt,
			formData["name"].(string),
			formData["address"].(string),
			formData["numberPhone"].(string),
			positionToInt,
		)

		if err != nil {
			log.Println(err)
			return c.JSON(fiber.Map{
				"error":  err.Error(),
				"status": fiber.StatusInternalServerError,
			})
		}

		return c.JSON(fiber.Map{
			"error":  nil,
			"status": c.Response().StatusCode(),
			"data":   formData,
			"result": newEmpResult,
		})
	})

	// Menambah position baru
	app.Post("/employee/position/new", func(c *fiber.Ctx) error {
		var formData map[string]any // variabel untuk menyimpan data yang diterima dari client-side

		err := c.BodyParser(&formData)
		if err != nil {
			log.Println(err)
			return c.JSON(fiber.Map{
				"error":  err.Error(),
				"status": fiber.StatusInternalServerError,
			})
		}

		if formData["name"].(string) == "" {
			return c.JSON(fiber.Map{
				"error":  "Position Filed is Empty!",
				"status": fiber.StatusInternalServerError,
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
				"status": fiber.StatusInternalServerError,
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
