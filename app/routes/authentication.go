package routes

import (
	"log"
	"logistica/app/models"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

// var store *session.Store = session.New()

func AuthenticationRoutes(app *fiber.App, store *session.Store) {
	app.Get("/login", func(c *fiber.Ctx) error {
		// session, _ := store.Get(c)

		// username := session.Get("username")
		// if username != nil {
		// 	return c.Redirect("/")
		// }
		return c.Render("login", fiber.Map{
			"Title": "LOGISTICA",
		})
	})

	app.Post("/login", func(c *fiber.Ctx) error {
		username := c.FormValue("username")
		password := c.FormValue("password")
		stayLoggedIn := c.FormValue("stay")
		var loggedIn bool = false

		users := models.User{}
		user := users.FindAll()[0]

		if user.Username != username {
			log.Println("Username Not Found")
			return c.Render("login", fiber.Map{
				"error": "Username Not Found.",
			})
		} else if user.Password != password {
			log.Println("Password Incorrect")
			return c.Render("login", fiber.Map{
				"error": "Password Incorrect.",
			})
		}
		session, err := store.Get(c)
		if err != nil {
			return c.Redirect("/505")
		}

		if stayLoggedIn == "true" {
			loggedIn = true
		}

		session.Set("username", user.Username)
		session.Set("user_id", user.Password)
		session.Set("logged_in", loggedIn)
		store.CookieHTTPOnly = true
		store.CookieSecure = true
		usernameSession := session.Get("username")

		if err := session.Save(); err != nil {
			log.Println(err)
			return err
		}

		log.Println("Login:", usernameSession)
		return c.Redirect("/dashboard")
	})
}

func DeauthenticationRoutes(app *fiber.App, store *session.Store) {
	app.Get("/logout", func(c *fiber.Ctx) error {
		session, err := store.Get(c)
		if err != nil {
			log.Println(err)
			return err
		}
		session.Delete("username")
		session.Destroy()
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
