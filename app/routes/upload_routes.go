package routes

import (
	"logistica/app/controllers"

	"github.com/gofiber/fiber/v2"
)

func FileManagement(app *fiber.App) {
	app.Post("/upload/image", controllers.UploadImage)
	app.Delete("/delete/image/:filename", controllers.DeleteImage)
}
