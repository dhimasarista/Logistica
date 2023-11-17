package routes

import (
	"logistica/app/controllers"

	"github.com/gofiber/fiber/v2"
)

func NotFoundError(app *fiber.App) {
	app.Get("/404", controllers.IndexRender)
}

func InternalServerError(app *fiber.App) {
	app.Get("/500", controllers.IndexRender)
}
