package routes

import (
	"encoding/json"
	"log"
	"logistica/app/controllers"
	"logistica/app/models"
	"logistica/app/utility"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

func InventoryRoutes(app *fiber.App, store *session.Store) {
	var productModel *models.Product = &models.Product{}
	var manufacturerModel *models.Manufacturer = &models.Manufacturer{}
	categoryModel := &models.Category{}

	// app.Put("/product/id")

	app.Delete("/product/:id", func(c *fiber.Ctx) error {
		// Mendapatkan ID produk dari parameter URL
		productID := c.Params("id")

		// Mengonversi ID produk ke dalam bentuk integer
		productIDInteger, err := strconv.Atoi(productID)
		if err != nil {
			return c.JSON(fiber.Map{
				"status": fiber.StatusBadRequest,
				"error":  err.Error(),
			})
		}

		// Mengambil produk dari model berdasarkan ID
		err = productModel.DeleteProduct(productIDInteger)
		if err != nil {
			return c.JSON(fiber.Map{
				"status": fiber.StatusBadRequest,
				"error":  err.Error(),
			})
		}

		// Mengembalikan respons JSON dengan data produk
		return c.JSON(fiber.Map{
			"error":  nil,
			"status": fiber.StatusOK,
		})
	})

	app.Post("/product/new", func(c *fiber.Ctx) error {
		// Mendapatkan ID terakhir dari produk
		lastId, err := productModel.LastId()
		if err != nil {
			log.Println(err)
			return c.JSON(fiber.Map{
				"error":  err.Error(),
				"status": fiber.StatusInternalServerError,
			})
		}

		var formData map[string]string // Variabel untuk menyimpan data yang diterima dari client-side
		body := c.Body()
		err = json.Unmarshal(body, &formData)
		if err != nil {
			log.Println(err)
			return c.JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		// Memeriksa apakah data penting seperti 'manufacturer' dan 'name' kosong
		if string(formData["manufacturer"]) == "" || string(formData["name"]) == "" || string(formData["category"]) == "" {
			return c.JSON(fiber.Map{
				"error":  "Form is Empty",
				"status": fiber.StatusBadRequest,
			})
		}

		// Memeriksa data manufacturer yang diterima
		var manufacturerData int
		// Jika dikirim dalam bentuk string number
		if utility.IsNumeric(formData["manufacturer"]) {
			// Data yang dikirim dalam bentuk string number adalah data yang sudah ada
			// Kemudian data dikonversi dari str ke integer
			manufacturerStrToInt, _ := strconv.Atoi(formData["manufacturer"])
			manufacturerData = manufacturerStrToInt
		} else {
			// Jika yang dikirim adalah string char, maka dibuat row data manufacturer baru
			// Dengan patokan id terakhir
			lastIdManufacturer, _ := manufacturerModel.LastId()
			var newIdManufacturer = lastIdManufacturer + 1
			_, err := manufacturerModel.NewManufacturer(newIdManufacturer, formData["manufacturer"])
			if err != nil {
				log.Println(err)
				return c.JSON(fiber.Map{
					"error": err.Error(),
				})
			}
			manufacturerData = newIdManufacturer
		}

		var categoryData int
		if utility.IsNumeric(formData["category"]) {
			categoryStrToInt, _ := strconv.Atoi(formData["category"])
			categoryData = categoryStrToInt
		} else {
			lastIdCategory, _ := categoryModel.LastId()
			var newIdCategory = lastIdCategory + 1
			_, err := categoryModel.NewCategory(newIdCategory, formData["category"])
			if err != nil {
				log.Println(err)
				return c.JSON(fiber.Map{
					"error": err.Error(),
				})
			}
			categoryData = newIdCategory
		}

		// Mengkonversi string ke integer untuk beberapa atribut
		stocksStrToInt, _ := strconv.Atoi(formData["stockAmount"])
		priceStrToInt, _ := strconv.Atoi(formData["price"])
		weightStrToint, _ := strconv.Atoi(formData["weight"])

		// Menyusun hasil data produk
		// results := map[string]interface{}{
		// 	"id":            lastId + 1,
		// 	"name":          string(formData["name"]),
		// 	"serial_number": string(formData["name"]),
		// 	"manufacturer":  manufacturerData,
		// 	"stocks":        stocksStrToInt,
		// 	"price":         priceStrToInt,
		// 	"weight":        weightStrToint,
		// 	"category":      categoryData,
		// }

		results, err := productModel.NewProduct(
			lastId+1,
			string(formData["name"]),
			string(formData["serialNumber"]),
			manufacturerData,
			stocksStrToInt,
			priceStrToInt,
			weightStrToint,
			categoryData,
		)

		if err != nil {
			log.Println(err)
			return c.JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		// Mengembalikan respons JSON setelah menambahkan produk baru
		return c.JSON(fiber.Map{
			"error":   nil,
			"status":  fiber.StatusOK,
			"message": "Success Add New Product",
			"result":  results,
		})
	})

	app.Get("/inventory", func(c *fiber.Ctx) error {
		// Mendapatkan path dari URL
		var path string = c.Path()

		// Mendapatkan nama pengguna dari sesi
		var username string = controllers.GetSessionUsername(c, store)

		// Mengambil semua produk dari model
		products, err := productModel.FindAll()
		if err != nil {
			log.Println(err)
			InternalServerError(c, err.Error())
		}

		// Mengambil semua manufaktur dari model
		manufacturers, err := manufacturerModel.FindAll()
		if err != nil {
			log.Println(err)
			InternalServerError(c, err.Error())
		}

		// Mengambil semua kategori dari model
		categories, err := categoryModel.FindAll()
		if err != nil {
			log.Println(err)
			InternalServerError(c, err.Error())
		}

		// Merender halaman "inventory_page" dengan data yang diperlukan
		return c.Render("inventory_page", fiber.Map{
			"path":          path,
			"user":          username,
			"products":      products,
			"manufacturers": manufacturers,
			"categories":    categories,
		})
	})
	// Endpoint untuk mendapatkan detail produk berdasarkan ID
	app.Get("/inventory/product/:id", func(c *fiber.Ctx) error {
		// Mendapatkan ID produk dari parameter URL
		productID := c.Params("id")

		// Mengonversi ID produk ke dalam bentuk integer
		productIDInteger, err := strconv.Atoi(productID)
		if err != nil {
			return c.JSON(fiber.Map{
				"status": fiber.StatusBadRequest,
				"error":  err.Error(),
			})
		}

		// Mengambil produk dari model berdasarkan ID
		err = productModel.GetById(productIDInteger)
		if err != nil {
			return c.JSON(fiber.Map{
				"status": fiber.StatusBadRequest,
				"error":  err.Error(),
			})
		}

		// Menyiapkan data untuk respons JSON
		data := map[string]interface{}{
			"id":             productModel.ID.Int64,
			"name":           productModel.Name.String,
			"categoryId":     productModel.CategoryID.Int64,
			"category":       productModel.CategoryName.String,
			"manufacturerId": productModel.ManufacturerID.Int64,
			"manufacturer":   productModel.ManufacturerName.String,
			"price":          productModel.Price.Int64,
			"serialNumber":   productModel.SerialNumber.String,
			"stocks":         productModel.Stocks.Int64,
			"weight":         productModel.Weight.Int64,
		}

		// Mengembalikan respons JSON dengan data produk
		return c.JSON(fiber.Map{
			"data": data,
		})
	})

	// Endpoint untuk memeriksa keberadaan ID dalam database
	app.Get("/inventory/check/:id", func(c *fiber.Ctx) error {
		// Mendapatkan ID dari parameter URL
		idString := c.Params("id")
		idInteger, err := strconv.Atoi(idString)
		var isIdExists bool = false

		// Menangani kesalahan jika konversi ID ke integer tidak berhasil
		if err != nil {
			log.Println(err)
		}

		// Jika ID adalah 0, mengindikasikan produk baru
		if idInteger == 0 {
			return c.JSON(fiber.Map{
				"new_product": true,
			})
		}

		// Memeriksa keberadaan ID dalam database
		dataId, err := productModel.OnlyGetID(idInteger)
		if err != nil {
			log.Println(err)
			return c.JSON(fiber.Map{
				"error":  err.Error(),
				"status": 500,
			})
		}

		// Menetapkan isIdExists menjadi true jika ID ditemukan dalam database
		if int64(dataId) == int64(idInteger) {
			isIdExists = true
		}

		// Mengembalikan respons JSON dengan hasil pemeriksaan ID
		return c.JSON(fiber.Map{
			"is_id_exists":    isIdExists,
			"response_status": c.Response().StatusCode(),
		})
	})

	// Endpoint untuk memperbarui stok produk berdasarkan ID
	app.Post("/inventory/stocks/update", func(c *fiber.Ctx) error {
		// Mendeklarasikan variabel untuk menyimpan data yang diterima dari client-side
		var formData map[string]interface{}

		// Memparsing data dari body request menjadi bentuk map
		err := c.BodyParser(&formData)
		if err != nil {
			log.Println("Body Parser", err)
			return c.JSON(fiber.Map{
				"error":  err.Error(),
				"status": fiber.StatusInternalServerError,
			})
		}

		if formData["id"].(string) == "" {
			return c.JSON(fiber.Map{
				"error":  "Form is Empty",
				"status": fiber.StatusBadRequest,
			})
		}

		// Mengambil ID produk dari formData dan mengonversinya ke dalam tipe data integer
		idStr := formData["id"].(string)
		id, _ := strconv.Atoi(idStr)

		// Mendapatkan stok terakhir dari produk
		lastStocks, err := productModel.LastStocks(id)
		if err != nil {
			log.Println(err)
			return c.JSON(fiber.Map{
				"error":  err.Error(),
				"status": fiber.StatusInternalServerError,
			})
		}

		// Mengambil jumlah stok yang akan ditambahkan dari formData dan mengonversinya ke dalam tipe data integer
		stockStr := formData["amountStocks"].(string)
		stock, _ := strconv.Atoi(stockStr)
		if stock < 0 {
			return c.JSON(fiber.Map{
				"error":  "Cannot be a Negative Number!",
				"status": fiber.StatusBadRequest,
			})
		}

		// Memperbarui stok produk
		results, err := productModel.UpdateStocks(id, lastStocks+stock)
		if err != nil {
			log.Println(err)
			return c.JSON(fiber.Map{
				"error":  err.Error(),
				"status": fiber.StatusInternalServerError,
			})
		}

		// Mengembalikan respons JSON dengan pesan sukses dan hasil pembaruan stok
		return c.JSON(fiber.Map{
			"message": "Stock Updated!",
			"results": results,
			"status":  c.Response().StatusCode(),
		})
	})
}

/*
	Server tidak dapat menerima atau merespons data kecuali jika data tersebut
	berbentuk formData["data"].(string) atau interface{}.
	- Jika formData["data"].(int), tidak akan ada respons error.
	- Data yang diterima dari klien dipaksa menjadi string sebelum dikonversi
	  sesuai dengan tipe data yang akan diterima oleh DBMS.
*/
