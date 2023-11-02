package routes

import (
	"logistica/app/controllers"

	"github.com/gofiber/fiber/v2"
)

func AuthRoutes(app *fiber.App) {
	app.Get("/login", controllers.LoginRender)
}
