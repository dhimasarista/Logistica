package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

func UserAuthorization(app *fiber.App) {
	store := session.New()

	app.Use(func(c *fiber.Ctx) error {
		session, _ := store.Get(c)

		authenticated := session.Get("authenticated").(bool)
		if !authenticated {
			return c.Redirect("login")
		}

		return c.Next()
	})
}
