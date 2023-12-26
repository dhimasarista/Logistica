package routes

import (
	"database/sql"
	"log"
	"logistica/app/config"
	"logistica/app/controllers"
	"logistica/app/models"
	"logistica/app/utility"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

func OrdersRoutes(app *fiber.App, store *session.Store) {
	var product = &models.Product{}
	var order = &models.Order{}
	var orderDetail = &models.OrderDetail{}
	var earning = &models.Earning{}
	var db *sql.DB

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
		time.Sleep(1 * time.Second)
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
		db = config.ConnectSQLDB()
		defer db.Close()
		tx, err := db.Begin()
		if err != nil {
			log.Println(err)
			return c.JSON(fiber.Map{
				"error":  err.Error(),
				"status": fiber.StatusInternalServerError,
			})
		}
		time.Sleep(1 * time.Second)
		var formData map[string]string
		err = c.BodyParser(&formData)
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

		var totalPrice int                                         // Variabel untuk menampung total price
		product.GetById(idProduct)                                 // Mengambil data price ke database
		totalPriceCalculate := &totalPrice                         // Nembak memory-address totalPrice
		*totalPriceCalculate = quantity * int(product.Price.Int64) // hasil disimpan di totalPrice

		if quantity == 0 {
			return c.JSON(fiber.Map{
				"error":  "Quantity is Zero",
				"status": fiber.StatusInternalServerError,
			})
		} else if quantity < 0 {
			return c.JSON(fiber.Map{
				"error":  "Bad Request",
				"status": fiber.StatusInternalServerError,
			})
		}

		// Create data order detail terlebih dahulu
		_, err = orderDetail.NewOrder(tx, newOrderID, buyer, numberPhone, address)
		if err != nil {
			log.Println(err)
			tx.Rollback()
			return c.JSON(fiber.Map{
				"error":  err.Error(),
				"status": fiber.StatusInternalServerError,
			})
		}
		// Jika tidak error, lanjut ke Order
		err = order.NewOrder(tx, int64(newOrderID), int64(quantity), int64(totalPrice), int64(idProduct), int64(orderStatusID), int64(newOrderID))
		if err != nil {
			log.Println(err)
			tx.Rollback()
			return c.JSON(fiber.Map{
				"error":  err.Error(),
				"status": fiber.StatusInternalServerError,
			})
		}

		// Kemudian memasukkan pendapatan ke tabel earnings
		err = earning.NewOrder(tx, newOrderID, totalPrice, product.ManufacturerName.String+" "+product.Name.String, quantity, int(product.Price.Int64))
		if err != nil {
			log.Println(err)
			tx.Rollback()
			return c.JSON(fiber.Map{
				"error":  err.Error(),
				"status": fiber.StatusInternalServerError,
			})
		}

		// Commit transaksi jika semua operasi berhasil
		tx.Commit()

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
