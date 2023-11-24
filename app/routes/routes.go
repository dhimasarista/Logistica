package routes

import (
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	IndexRoutes(app)
	AuthRoutes(app)
	DashboardRoutes(app)
	UnggahSuratRoutes(app)
	OrdersRoutes(app)
	InventoryRoutes(app)
	NotFoundError(app)
	InternalServerError(app)
}
