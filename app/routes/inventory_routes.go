package routes

import (
	"database/sql"
	"fmt"
	"log"
	"logistica/app/controllers"
	"logistica/app/models"
	"reflect"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

func InventoryRoutes(app *fiber.App, store *session.Store) {
	var product *models.Product = &models.Product{}
	var manufacturer *models.Manufacturer = &models.Manufacturer{}
	category := &models.Category{}

	app.Get("/inventory", func(c *fiber.Ctx) error {
		var path string = c.Path()
		var username string = controllers.GetSessionUsername(c, store)

		products, err := product.FindAll()
		if err != nil {
			log.Println(err)
			InternalServerError(c, err.Error())
		}

		manufacturers, err := manufacturer.FindAll()
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

		fmt.Println(productID)
		productIDInteger, err := strconv.Atoi(productID)
		if err != nil {
			return c.JSON(fiber.Map{
				"status": fiber.StatusBadRequest,
				"error":  err.Error(),
			})
		}

		err = product.GetById(productIDInteger)
		if err != nil {
			return c.JSON(fiber.Map{
				"status": fiber.StatusBadRequest,
				"error":  err.Error(),
			})
		}

		data := map[string]interface{}{
			"id":             product.ID.Int64,
			"name":           product.Name.String,
			"categoryId":     product.CategoryID.Int64,
			"category":       product.CategoryName.String,
			"manufacturerId": product.ManufacturerID.Int64,
			"manufacturer":   product.ManufacturerName.String,
			"price":          product.Price.Int64,
			"serialNumber":   product.SerialNumber.String,
			"stocks":         product.Stocks.Int64,
			"weight":         product.Weight.Int64,
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

		dataId, err := product.OnlyGetID(idInteger)
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
		fmt.Println(reflect.TypeOf(formData["id"]))

		idStr := formData["id"].(string)
		id, _ := strconv.Atoi(idStr)
		lastStocks, err := product.LastStocks(id)
		if err != nil {
			log.Println(err)
			return c.JSON(fiber.Map{
				"error":  err.Error(),
				"status": fiber.StatusInternalServerError,
			})
		}

		stockStr := formData["amountStocks"].(string)
		stock, _ := strconv.Atoi(stockStr)
		results, err := product.UpdateStocks(id, lastStocks+stock)
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

	app.Post("/product/new", func(c *fiber.Ctx) error {
		lastId, err := product.LastId()
		if err != nil {
			log.Println(err)
			return c.JSON(fiber.Map{
				"error":  err.Error(),
				"status": fiber.StatusInternalServerError,
			})
		}
		if lastId <= 1020 {
			lastId = 1020
		}

		var formData map[string]interface{} // variabel untuk menyimpan data yang diterima dari client-side
		if err := c.BodyParser(&formData); err != nil {
			log.Println(err)
		}
		productData := &models.Product{
			ID:             sql.NullInt64{Int64: formData["id"].(int64) + 1},
			Name:           sql.NullString{String: formData["name"].(string)},
			SerialNumber:   sql.NullString{String: formData["serialNumber"].(string)},
			ManufacturerID: sql.NullInt64{Int64: formData["manufacturer"].(int64)},
			Stocks:         sql.NullInt64{Int64: formData["stocks"].(int64)},
			Price:          sql.NullInt64{Int64: formData["price"].(int64)},
			Weight:         sql.NullInt64{Int64: formData["weight"].(int64)},
			CategoryID:     sql.NullInt64{Int64: formData["category"].(int64)},
		}

		fmt.Println(productData)

		// results, err := productData.NewProduct()
		// if err != nil {
		// 	return c.JSON(fiber.Map{
		// 		"error":  err.Error(),
		// 		"status": fiber.StatusInternalServerError,
		// 	})
		// }

		return c.JSON(fiber.Map{
			"error":  nil,
			"status": c.Response().StatusCode(),
			"data":   productData,
			"result": nil,
		})
	})
}
