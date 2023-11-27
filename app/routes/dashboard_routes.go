package routes

import (
	"fmt"
	"log"

	"github.com/go-resty/resty/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

func DashboardRoutes(app *fiber.App, store *session.Store, client *resty.Client) {
	app.Get("/dashboard", func(c *fiber.Ctx) error {
		var path string = c.Path()
		session, err := store.Get(c)
		if err != nil {
			log.Println(err)
			return err
		}
		username := session.Get("username")
		defer session.Save()

		metrics, err := client.R().Get("/metrics")
		if err != nil {
			log.Println(err)
		}

		fmt.Println(metrics.Result())

		// Mengirimkan halaman HTML yang dihasilkan ke browser
		return c.Render("dashboard", fiber.Map{
			"path": path,
			"user": username,
		})
	},
	)
}
