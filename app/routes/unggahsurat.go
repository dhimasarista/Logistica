package routes

import (
	"logistica/app/controllers"

	"github.com/gofiber/fiber/v2"
)

func UnggahSuratRoutes(app *fiber.App) {
	app.Get("/unggahsurat", controllers.UnggahSuratRender)
}
