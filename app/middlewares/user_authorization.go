package middlewares

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

func UserAuthorization(app *fiber.App, store *session.Store) {
	app.Use(func(c *fiber.Ctx) error {
		session, err := store.Get(c)
		if err != nil {
			return err
		}
		username := session.Get("username")
		session.Save()

		// Check if the user is not logged in
		if username != nil {
			fmt.Println("UserAuth != nil:", username)
			return c.Next()
		}
		fmt.Println("UserAuth:", username)

		return c.Redirect("/login")
	})
}
