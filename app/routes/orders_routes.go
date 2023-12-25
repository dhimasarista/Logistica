package routes

import (
	"log"
	"logistica/app/controllers"
	"logistica/app/models"
	"logistica/app/utility"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

func OrdersRoutes(app *fiber.App, store *session.Store) {
	var product = &models.Product{}
	var order = &models.Order{}
	var orderDetail = &models.OrderDetail{}

	app.Get("/orders", func(c *fiber.Ctx) error {
		var path string = c.Path()
		var username string = controllers.GetSessionUsername(c, store)

		products, err := product.FindAll()
		if err != nil {
			log.Println(err)
			return InternalServerError(c, err.Error())
		}

		orders, err := order.FindAll()
		if err != nil {
			log.Println(err)
			return InternalServerError(c, err.Error())
		}

		return c.Render("orders_page", fiber.Map{
			"path":     path,
			"user":     username,
			"products": products,
			"orders":   orders,
		})
	})

	app.Post("/order/calculate", func(c *fiber.Ctx) error {
		var formData map[string]string
		// Mengambil data dari body yang dikirim oleh client
		err := c.BodyParser(&formData)
		if err != nil {
			log.Println(err)
			return c.JSON(fiber.Map{
				"error":  err.Error(),
				"status": fiber.StatusInternalServerError,
			})
		}
		if formData["quantity"] == "" || formData["idProduct"] == "" {
			return c.JSON(fiber.Map{
				"error":  "Form is Empty!",
				"status": fiber.StatusInternalServerError,
			})
		}

		quantityInt, _ := strconv.Atoi(formData["quantity"])
		productInt, _ := strconv.Atoi(formData["idProduct"])
		err = product.GetById(productInt)
		if err != nil {
			log.Println(err)
			return c.JSON(fiber.Map{
				"error":  err.Error(),
				"status": fiber.StatusInternalServerError,
			})
		}

		if quantityInt > int(product.Stocks.Int64) {
			return c.JSON(fiber.Map{
				"error":  "Stock is Not Enough.",
				"status": fiber.StatusBadRequest,
			})
		}

		data := map[string]any{
			"id":      product.ID.Int64,
			"name":    product.Name.String,
			"price":   utility.RupiahFormat(product.Price.Int64),
			"payment": utility.RupiahFormat(int64(quantityInt) * product.Price.Int64),
		}

		return c.JSON(fiber.Map{
			"error":  nil,
			"status": c.Response().StatusCode(),
			"data":   data,
		})
	})

	app.Post("/order/new", func(c *fiber.Ctx) error {
		var formData map[string]string
		err := c.BodyParser(&formData)
		if err != nil {
			log.Println(err)
			return c.JSON(fiber.Map{
				"error":  err.Error(),
				"status": fiber.StatusInternalServerError,
			})
		}

		// Mengambil Last ID untuk Orders dan Order Detail
		// Dengan patokan dari last ID order_detail
		lastIdOrderDetail, err := orderDetail.LastId()
		newOrderID := lastIdOrderDetail + 1
		if err != nil {
			log.Println(err)
			return c.JSON(fiber.Map{
				"error":  err.Error(),
				"status": fiber.StatusInternalServerError,
			})
		}

		// OrderData
		idProduct, _ := strconv.Atoi(formData["idProduct"])
		buyer := formData["buyer"]
		numberPhone := formData["numberPhone"]
		address := formData["address"]
		quantity, _ := strconv.Atoi(formData["quantity"])
		orderStatusID := 1

		// Create data order detail terlebih dahulu
		_, err = orderDetail.NewOrder(newOrderID, buyer, numberPhone, address)
		if err != nil {
			log.Println(err)
			return c.JSON(fiber.Map{
				"error":  err.Error(),
				"status": fiber.StatusInternalServerError,
			})
		}
		// Jika tidak error, lanjut ke Order
		err = order.NewOrder(int64(newOrderID), int64(quantity), int64(idProduct), int64(orderStatusID), int64(newOrderID))
		if err != nil {
			log.Println(err)
			return c.JSON(fiber.Map{
				"error":  err.Error(),
				"status": fiber.StatusInternalServerError,
			})
		}

		// Ketika order masuk maka stok akan dikurangin
		// Mengambil stok terakhir terlebih dahulu
		lastStock, _ := product.LastStocks(idProduct)
		product.UpdateStocks(idProduct, lastStock-quantity)

		return c.JSON(fiber.Map{
			"error":   nil,
			"status":  c.Response().StatusCode(),
			"data":    formData,
			"message": "Success Add Order",
		})
	})
}
