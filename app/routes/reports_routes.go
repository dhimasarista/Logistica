package routes

import (
	"log"
	"logistica/app/controllers"
	"logistica/app/models"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

func ReportsRoutes(app *fiber.App, store *session.Store) {
	var stockRecord = &models.StockRecord{}
	app.Get("/reports", func(c *fiber.Ctx) error {
		// Mendapatkan path dan username dari sesi
		var path string = c.Path()
		var username string = controllers.GetSessionUsername(c, store)

		stockRecords, err := stockRecord.FindAll()
		if err != nil {
			log.Println(err)
			InternalServerError(c, err.Error())
		}

		return c.Render("reports_page", fiber.Map{
			"path":          path,
			"user":          username,
			"stock_records": stockRecords,
		})
	})
}
