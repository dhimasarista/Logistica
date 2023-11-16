package middlewares

import (
	"github.com/gofiber/fiber/v2"
)

func UserAuthorization(app *fiber.App) {
	// store := session.New()

	// app.Use(func(c *fiber.Ctx) error {
	// 	session, _ := store.Get(c)

	// 	log.Println(session.ID())

	// 	sessionID := session.ID()

	// 	if sessionID == "" {
	// 		c.Redirect("/login")
	// 	}

	// 	return c.Next()
	// })

	// app.Use(basicauth.New(basicauth.Config{
	// 	Users: map[string]string{
	// 		"john":  "doe",
	// 		"admin": "123456",
	// 	},
	// 	Realm: "Forbidden",
	// 	Authorizer: func(user, pass string) bool {
	// 		if user == "john" && pass == "doe" {
	// 			return true
	// 		}
	// 		if user == "admin" && pass == "123456" {
	// 			return true
	// 		}
	// 		return false
	// 	},
	// 	Unauthorized: func(c *fiber.Ctx) error {
	// 		return c.Redirect("/login")
	// 	},
	// 	ContextUsername: "_user",
	// 	ContextPassword: "_pass",
	// }))
}
