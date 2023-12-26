package routes

import (
	"database/sql"
	"fmt"
	"log"
	"strings"

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
	ReportsRoutes(app, store)
	FileManagement(app)
}

// Routing tambahan
func InternalServerError(c *fiber.Ctx, message string) error {
	log.Println(message)
	var messageFormatted = strings.Replace(message, " ", "+", -1)
	var path = fmt.Sprintf("/error?code=500&title=Internal+Server+Error&message=%s", messageFormatted)
	return c.Redirect(path)
}

// Fungsi-fungsi handler
func HandleErrorAndRollback(tx *sql.Tx, err error, c *fiber.Ctx) error {
	if err != nil {
		log.Println(err)
		tx.Rollback()
		return c.JSON(fiber.Map{
			"error":  err.Error(),
			"status": fiber.StatusInternalServerError,
		})
	}
	return nil
}
