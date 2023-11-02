package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/internal/storage/memory"
	"github.com/gofiber/fiber/v2/middleware/session"
)

func UserAuthorization(app *fiber.App) {
	store := memory.New()
	sess := session.New(session.Config{
		Storage: store,
	})
	app.Use(func(c *fiber.Ctx) error {
		session, _ := sess.Get(c)

		authenticated, ok := session.Get("authenticated").(bool)
		if !ok || !authenticated {
			return c.Redirect("login")
		}

		return c.Next()
	})
}
