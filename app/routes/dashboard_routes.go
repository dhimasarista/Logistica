package routes

import (
	"logistica/app/controllers"

	"github.com/gofiber/fiber/v2"
)

func DashboardRoutes(app *fiber.App) {
	app.Get("/dashboard", controllers.DashboardRender)
}
