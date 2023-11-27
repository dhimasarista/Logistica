package routes

import (
	"github.com/go-resty/resty/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

func SetupRoutes(app *fiber.App, store *session.Store, client *resty.Client) {
	IndexRoutes(app)
	AuthenticationRoutes(app, store)
	DeauthenticationRoutes(app, store)
	DashboardRoutes(app, store, client)
	UnggahSuratRoutes(app)
	OrdersRoutes(app)
	InventoryRoutes(app)
	NotFoundError(app)
	InternalServerError(app)
}
