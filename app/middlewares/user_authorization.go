package middlewares

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

func UserAuthorization(app *fiber.App) {
	store := session.New()

	app.Use(func(c *fiber.Ctx) error {
		session, _ := store.Get(c)

		authenticated := session.Get("authenticated")
		log.Println(authenticated)

		return c.Next()
	})
}
