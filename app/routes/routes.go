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
	DashboardRoutes(app, store)
	OrdersRoutes(app, store)
	InventoryRoutes(app, store)
	ErrorRoutes(app, store)
	EmployeesRoutes(app, store)
	UploadFileRoutes(app)
}
