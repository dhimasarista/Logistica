package routes

import (
	"logistica/app/controllers"

	"github.com/gofiber/fiber/v2"
)

func InventoryRoutes(app *fiber.App) {
	app.Get("/inventory", controllers.InventoryRender)
}
