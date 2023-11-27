package routes

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

func DashboardRoutes(app *fiber.App, store *session.Store) {
	app.Get("/dashboard", func(c *fiber.Ctx) error {
		var path string = c.Path()
		session, err := store.Get(c)
		if err != nil {
			log.Println(err)
			return err
		}
		username := session.Get("username")
		defer session.Save()

		// Mengirimkan halaman HTML yang dihasilkan ke browser
		return c.Render("dashboard", fiber.Map{
			"path": path,
			"user": username,
		})
	},
	)
}
