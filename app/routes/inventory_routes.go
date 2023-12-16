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
}
