package middlewares

import (
	"github.com/gofiber/fiber/v2"
)

func UserAuthorization(app *fiber.App) {
	app.Use("/", func(c *fiber.Ctx) error {
		if c.Cookies("user") == "" {
			return c.Redirect("/login")
		}
		return c.Next()
	})
}
