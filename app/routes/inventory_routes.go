package routes

import (
	"log"
	"logistica/app/controllers"
	"logistica/app/models"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

func InventoryRoutes(app *fiber.App, store *session.Store) {
	app.Get("/inventory", func(c *fiber.Ctx) error {
		var path string = c.Path()
		var username string = controllers.GetSessionUsername(c, store)

		// Create a new instance of the Product model for each request
		product := &models.Product{}

		products, err := product.FindAll()
		if err != nil {
			log.Println(err)
			InternalServerError(c, err.Error())
		}

		return c.Render("inventory_page", fiber.Map{
			"path":     path,
			"user":     username,
			"products": products,
		})
	})
}
