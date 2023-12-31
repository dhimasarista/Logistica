package routes

import (
	"database/sql"
	"encoding/json"
	"log"
	"logistica/app/controllers"
	"logistica/app/helpers"
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
	position := &models.Position{}

	// Merender halaman employees_page
	app.Get("/employees", func(c *fiber.Ctx) error {
		// Mendapatkan path dan username dari sesi
		var path string = c.Path()
		var username string = controllers.GetSessionUsername(c, store)
		// Mengambil data employee dan posisi dari database
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
		// Mengembalikan respons dalam bentuk page yang dirender
		// Dan data juga dikirim dalam bentuk rendered-data
		return c.Render("employees_page", fiber.Map{
			"path":      path,
			"user":      username,
			"employees": employees,
			"positions": positions,
			"status":    c.Response().StatusCode(),
		})
	})
	// Mengambil data employee berdasarkan id
	app.Get("/employee/:id", func(c *fiber.Ctx) error {
		idStr := c.Params("id")
		idInteger, _ := strconv.Atoi(idStr)

		err := employee.GetById(int64(idInteger))
		if err != nil {
			log.Println(err)
			return c.JSON(fiber.Map{
				"error":  err.Error(),
				"status": 500,
			})
		}

		return c.JSON(fiber.Map{
			"data":   *employee,
			"status": c.Response().StatusCode(),
		})
	})
	// Membuat data employee baru
	app.Post("/employee/new", func(c *fiber.Ctx) error {

		// variabel untuk menyimpan data yang diterima dari client-side
		var formData map[string]interface{}
		err := c.BodyParser(&formData)
		if err != nil {
			log.Println(err)
			return c.JSON(fiber.Map{
				"error":  err.Error(),
				"status": fiber.StatusInternalServerError,
			})
		}
		// Memeriksa data yang dikirim
		isEmpty := helpers.IsEmpty(c, formData, "id", "name", "numberPhone", "position")
		if isEmpty {
			// Jika kosong `Empty` kirim response `bad request`
			return c.JSON(fiber.Map{
				"error":  "Form is Empty!",
				"status": fiber.StatusBadRequest,
			})
		}
		// Data yang dikirim dari client dalam bentuk string
		// Golang tidak akan menerima tipe data any mentah-mentah
		// Tidak akan ada error, tapi program tetap anomali
		// Kemudian data dikonversi ke tipe data masing-masing
		idToInt, err := strconv.Atoi(formData["id"].(string))
		if err != nil {
			panic(err)
		}
		positionToInt, err := strconv.Atoi(formData["position"].(string))
		if err != nil {
			panic(err)
		}

		// Data yang diterima tadi langsung dieksekusi oleh basis data
		employee = &models.Employee{
			ID:          sql.NullInt64{Int64: int64(idToInt)},
			Name:        sql.NullString{String: formData["name"].(string)},
			Address:     sql.NullString{String: formData["address"].(string)},
			NumberPhone: sql.NullString{String: formData["numberPhone"].(string)},
			PositionID:  sql.NullInt64{Int64: int64(positionToInt)},
		}
		err = employee.NewEmployee()
		if err != nil {
			log.Println(err)
			return c.JSON(fiber.Map{
				"error":  err.Error(),
				"status": fiber.StatusInternalServerError,
			})
		}

		// Berhasil, kirimkan response 200 ke client
		return c.JSON(fiber.Map{
			"error":  nil,
			"status": c.Response().StatusCode(),
		})
	})
	// Memperbarui data employee
	app.Put("/employee", func(c *fiber.Ctx) error {
		var formData map[string]interface{} // Variabel untuk menyimpan data yang diterima dari client-side
		body := c.Body()
		err := json.Unmarshal(body, &formData)
		if err != nil {
			log.Println(err)
			return c.JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		isEmpty := helpers.IsEmpty(c, formData, "id", "name", "numberPhone", "position")
		if isEmpty {
			// Jika kosong `Empty` kirim response `bad request`
			return c.JSON(fiber.Map{
				"error":  "Form is Empty!",
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
		employee = &models.Employee{
			ID:          sql.NullInt64{Int64: int64(idToInt)},
			Name:        sql.NullString{String: formData["name"].(string)},
			Address:     sql.NullString{String: formData["address"].(string)},
			NumberPhone: sql.NullString{String: formData["numberPhone"].(string)},
			PositionID:  sql.NullInt64{Int64: int64(positionToInt)},
		}
		err = employee.UpdateEmployee()
		if err != nil {
			log.Println(err)
			return c.JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		return c.JSON(fiber.Map{
			"data":    formData,
			"message": "Sukses Update Employe!",
			"status":  c.Response().StatusCode(),
		})
	})
	// Menghapus employee berdasarkan id
	app.Delete("/employee/:id", func(c *fiber.Ctx) error {
		var id string = c.Params("id")        // Mengambil id Dari parameter
		var idInteger, err = strconv.Atoi(id) // Mengonversi dari string ke integer
		if err != nil {
			log.Println(err)
		}
		// Kemudian meneruskan ke basis data
		err = employee.DeleteEmployee(int64(idInteger))
		if err != nil {
			log.Println(err)
			return c.JSON(fiber.Map{
				"error":  err.Error(),
				"status": 500,
			})
		}
		// Jika berhasil, kirim response
		return c.JSON(fiber.Map{
			"status":  c.Response().StatusCode(), // 200 OK
			"message": "Employee Deleted!",
		})

	})
	// Memeriksa ketersedian ID
	app.Get("/employee/check/:id", func(c *fiber.Ctx) error {
		var id string = c.Params("id")     // Mendapatkan id dari paramater path
		idInteger, err := strconv.Atoi(id) // Konversi string ke integer
		if err != nil {
			log.Println(err)
		}
		// Mengambil ID dari basis data
		employee.GetById(int64(idInteger))
		// Menentukan status ketersedian id
		var isIdExists bool = false
		if employee.ID.Int64 == int64(idInteger) {
			isIdExists = true
		}
		// Mengirimnya dalam bentuk JSON
		return c.JSON(fiber.Map{
			"is_exists": isIdExists,
			"status":    c.Response().StatusCode(), // 200 OK
		})
	})
	// Mengirim ID baru
	app.Get("/employee/new/id", func(c *fiber.Ctx) error {
		lastId, err := employee.LastId() // Mengambil Max ID dari basis data
		if err != nil {
			log.Println(err)
			InternalServerError(c, err.Error())
		}

		if lastId <= 100020 {
			lastId = 100020
		}
		// Berhasil, kirim id terakhir yang diminta
		return c.JSON(fiber.Map{
			"newId":  lastId + 1,
			"status": c.Response().StatusCode(),
		})
	})
	// Menambah position baru: Not Used and Fixed
	app.Post("/employee/position/new", func(c *fiber.Ctx) error {
		// variabel untuk menyimpan data yang diterima dari client-side
		var formData map[string]any
		// Mengambil data dari body yang dikirim oleh client
		err := c.BodyParser(&formData)
		if err != nil {
			log.Println(err)
			return c.JSON(fiber.Map{
				"error":  err.Error(),
				"status": fiber.StatusInternalServerError,
			})
		}
		// Memeriksa kekosongan data yang dikirim
		if formData["name"].(string) == "" {
			return c.JSON(fiber.Map{
				"error":  "Position Filed is Empty!",
				"status": fiber.StatusInternalServerError,
			})
		}
		// Mengambil id
		lastId, err := position.LastId()
		if err != nil {
			log.Println(err)
			return err
		}

		position = &models.Position{
			ID:   sql.NullInt64{Int64: int64(lastId + 1)},
			Name: sql.NullString{String: formData["name"].(string)},
		}
		err = position.NewPosition()
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
		})
	})

}
