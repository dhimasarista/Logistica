package routes

import (
	"database/sql"
	"fmt"
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
	var earning = &models.Earning{}
	var stockRecord = &models.StockRecord{}
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

	app.Get("/order/detail/:id", func(c *fiber.Ctx) error {
		var id = c.Params("id")
		idAtoi, _ := strconv.Atoi(id)

		err := order.GetByID(idAtoi)
		if err != nil {
			log.Println(err)
			return c.JSON(fiber.Map{
				"error":  err.Error(),
				"status": fiber.StatusInternalServerError,
			})
		}
		var buttonDetail string
		var buttonHashedCode string
		var buttonCancelHashed string
		if order.Status.Name.String == "on process" {
			onDelivery, _ := utility.GenerateHash("on delivery")
			buttonHashedCode = onDelivery
			cancelOrder, _ := utility.GenerateHash("cancelled")
			buttonCancelHashed = cancelOrder
			buttonDetail = fmt.Sprintf(`<button type="button" class="btn btn-primary" onclick="processOrder('%d' ,'%s')">Ship</button>
			<button type="button" class="btn btn-danger" onclick="processOrder('%d', '%s')">Cancel</button>`, idAtoi, buttonHashedCode, idAtoi, buttonCancelHashed)
		} else if order.Status.Name.String == "on delivery" {
			received, _ := utility.GenerateHash("received")
			buttonHashedCode = received
			buttonDetail = fmt.Sprintf(`<button type="button" class="btn btn-success" processOrder('%d', '%s')>Finish</button>
			<button type="button" class="btn btn-danger">Others</button>`, idAtoi, buttonHashedCode)
		}

		return c.JSON(fiber.Map{
			"error":         nil,
			"status":        c.Response().StatusCode(),
			"data":          order,
			"button_detail": buttonDetail,
		})
	})

	app.Post("/order/next", func(c *fiber.Ctx) error {
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
		var formData map[string]string
		err = c.BodyParser(&formData)
		if err != nil {
			log.Println(err)
			return c.JSON(fiber.Map{
				"error":  err.Error(),
				"status": fiber.StatusInternalServerError,
			})
		}

		orderId, _ := strconv.Atoi(formData["order_id"])

		if utility.ValidateTextHashed(formData["status"], "on delivery") {
			err = order.UpdateOrder(orderId, 2)
			if err != nil {
				log.Println(err)
				return c.JSON(fiber.Map{
					"error":  err.Error(),
					"status": fiber.StatusInternalServerError,
				})
				// Kemudian memasukkan pendapatan ke tabel earnings
			}
			order.GetByID(orderId)
			err = earning.NewOrder(tx, int(order.TotalPrice.Int64), order.Product.Name.String, int(order.Pieces.Int64), int(order.Product.Price.Int64))
			if err != nil {
				log.Println(err)
				tx.Rollback()
				return c.JSON(fiber.Map{
					"error":  err.Error(),
					"status": fiber.StatusInternalServerError,
				})
			}
			tx.Commit()
			return c.JSON(fiber.Map{
				"error":   nil,
				"status":  c.Response().StatusCode(),
				"message": "Order on Delivery",
			})
		}

		return c.JSON(fiber.Map{
			"error":  "Internal Server Error",
			"status": fiber.StatusInternalServerError,
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

	// app.Post("/order/update

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

		// Jika tidak error, lanjut ke Order
		err = order.NewOrder(tx, buyer, numberPhone, address, int64(quantity), int64(totalPrice), int64(idProduct), int64(orderStatusID))
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

		// Mengirim data ke stockRecord
		stockRecord = &models.StockRecord{
			Amount:      sql.NullInt64{Int64: int64(quantity)},
			Before:      sql.NullInt64{Int64: int64(lastStock)},
			After:       sql.NullInt64{Int64: int64(lastStock - quantity)},
			IsAddition:  sql.NullBool{Bool: false},
			ProductID:   sql.NullInt64{Int64: int64(idProduct)},
			Description: sql.NullString{String: fmt.Sprintf("Order by %s", buyer)},
		}
		err = stockRecord.NewRecord()
		if err != nil {
			log.Println(err)
			return c.JSON(fiber.Map{
				"error":  err.Error(),
				"status": fiber.StatusInternalServerError,
			})
		}

		return c.JSON(fiber.Map{
			"error":   nil,
			"status":  c.Response().StatusCode(),
			"data":    formData,
			"message": "Success Add Order",
		})
	})
}
