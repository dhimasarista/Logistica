package routes

import (
	"fmt"
	"log"
	"logistica/app/models"
	"time"

	"github.com/gofiber/fiber/v2"
)

func AuthenticationRoutes(app *fiber.App) {

	app.Post("/login", func(c *fiber.Ctx) error {
		username := c.FormValue("username")
		password := c.FormValue("password")
		stayLoggedIn := c.FormValue("stay")
		var loggedIn time.Time = time.Now().Add(time.Minute * 60)

		users := models.User{}
		user := users.FindAll()[0]

		if user.Username != username {
			log.Println("Username Not Found")
			return c.Render("login", fiber.Map{
				"errors": []fiber.Map{
					{
						"message": "Username Not Found.",
					},
				},
			})
		} else if user.Password != password {
			log.Println("Password Incorrect")
			return c.Render("login", fiber.Map{
				"errors": []fiber.Map{
					{
						"message": "Password Incorrect.",
					},
				},
			})
		}

		if stayLoggedIn == "true" {
			loggedIn = time.Now().Add(time.Hour * 24 * 7)
		}

		c.Cookie(&fiber.Cookie{
			Name:     "user",
			Value:    user.Username,
			Expires:  loggedIn,
			HTTPOnly: true,
			Secure:   true,
		})
		fmt.Println("Berhasil Login")

		return c.Redirect("/dashboard")
	})

}

func DeauthenticationRoutes(app *fiber.App) {
	app.Get("/logout", func(c *fiber.Ctx) error {
		c.ClearCookie()
		return c.Redirect("/login")
	})
}

// func IsAuthenticated(c *fiber.Ctx) error {
// 	session, err := store.Get(c)
// 	if err != nil {
// 		return c.Redirect("/500")
// 	}

// 	username := session.Get("Username")
// 	if username == "" {

// 	}
// }
