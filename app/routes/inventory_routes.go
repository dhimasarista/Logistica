package routes

import (
	"encoding/json"
	"fmt"
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
	category := &models.Category{}

	app.Post("/product/new", func(c *fiber.Ctx) error {
		lastId, err := productModel.LastId()
		if err != nil {
			log.Println(err)
			return c.JSON(fiber.Map{
				"error":  err.Error(),
				"status": fiber.StatusInternalServerError,
			})
		}

		var formData map[string]string // variabel untuk menyimpan data yang diterima dari client-side
		body := c.Body()
		err = json.Unmarshal(body, &formData)
		if err != nil {
			log.Println(err)
			return c.JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		if string(formData["manufacturer"]) == "" || string(formData["name"]) == "" {
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
			// Jika yang dikirim adalah string char
		} else {
			// Maka dibuat row data manufacturer baru
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

		stocksStrToInt, _ := strconv.Atoi(formData["stocks"])
		priceStrToInt, _ := strconv.Atoi(formData["price"])
		weightStrToint, _ := strconv.Atoi(formData["weight"])
		categoryStrToInt, _ := strconv.Atoi(formData["category"])

		results := map[string]interface{}{
			"id":            lastId + 1,
			"name":          string(formData["name"]),
			"serial_number": string(formData["serialNumber"]),
			"manufacturer":  manufacturerData,
			"stocks":        stocksStrToInt,
			"price":         priceStrToInt,
			"weight":        weightStrToint,
			"category":      categoryStrToInt,
		}
		fmt.Println("After", formData)
		// fmt.Println("After", results)

		return c.JSON(fiber.Map{
			"error":   nil,
			"status":  fiber.StatusOK,
			"message": "Succes Add New Product",
			"result":  results,
		})
	})

	app.Get("/inventory", func(c *fiber.Ctx) error {
		var path string = c.Path()
		var username string = controllers.GetSessionUsername(c, store)

		products, err := productModel.FindAll()
		if err != nil {
			log.Println(err)
			InternalServerError(c, err.Error())
		}

		manufacturers, err := manufacturerModel.FindAll()
		if err != nil {
			log.Println(err)
			InternalServerError(c, err.Error())
		}

		categories, err := category.FindAll()
		if err != nil {
			log.Println(err)
			InternalServerError(c, err.Error())
		}

		return c.Render("inventory_page", fiber.Map{
			"path":          path,
			"user":          username,
			"products":      products,
			"manufacturers": manufacturers,
			"categories":    categories,
		})
	})

	// app.Post("/inventory/")

	app.Get("/inventory/product/:id", func(c *fiber.Ctx) error {
		productID := c.Params("id")

		productIDInteger, err := strconv.Atoi(productID)
		if err != nil {
			return c.JSON(fiber.Map{
				"status": fiber.StatusBadRequest,
				"error":  err.Error(),
			})
		}

		err = productModel.GetById(productIDInteger)
		if err != nil {
			return c.JSON(fiber.Map{
				"status": fiber.StatusBadRequest,
				"error":  err.Error(),
			})
		}

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

		return c.JSON(fiber.Map{
			"data": data,
		})
	})

	app.Get("/inventory/check/:id", func(c *fiber.Ctx) error {
		idString := c.Params("id")
		idInteger, err := strconv.Atoi(idString)
		var isIdExists bool = false
		if err != nil {
			log.Println(err)
		}
		if idInteger == 0 {
			return c.JSON(fiber.Map{
				"new_product": true,
			})
		}

		dataId, err := productModel.OnlyGetID(idInteger)
		if err != nil {
			log.Println(err)
			return c.JSON(fiber.Map{
				"error":  err.Error(),
				"status": 500,
			})
		}
		if int64(dataId) == int64(idInteger) {
			isIdExists = true
		}

		return c.JSON(fiber.Map{
			"is_id_exists":    isIdExists,
			"response_status": c.Response().StatusCode(),
		})
	})

	app.Post("/inventory/stocks/update", func(c *fiber.Ctx) error {
		var formData map[string]any
		err := c.BodyParser(&formData)
		if err != nil {
			log.Println("Body Parser", err)
			return c.JSON(fiber.Map{
				"error":  err.Error(),
				"status": fiber.StatusInternalServerError,
			})
		}

		idStr := formData["id"].(string)
		id, _ := strconv.Atoi(idStr)
		lastStocks, err := productModel.LastStocks(id)
		if err != nil {
			log.Println(err)
			return c.JSON(fiber.Map{
				"error":  err.Error(),
				"status": fiber.StatusInternalServerError,
			})
		}

		stockStr := formData["amountStocks"].(string)
		stock, _ := strconv.Atoi(stockStr)
		results, err := productModel.UpdateStocks(id, lastStocks+stock)
		if err != nil {
			log.Println(err)
			return c.JSON(fiber.Map{
				"error":  err.Error(),
				"status": fiber.StatusInternalServerError,
			})
		}

		return c.JSON(fiber.Map{
			"message": "Stock Updated!",
			"results": results,
			"status":  c.Response().StatusCode(),
		})
	})
}
