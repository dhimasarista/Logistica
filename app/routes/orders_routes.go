package routes

import (
	"logistica/app/controllers"

	"github.com/gofiber/fiber/v2"
)

func OrdersRoutes(app *fiber.App) {
	app.Get("/orders", controllers.OrdersRender)
}
