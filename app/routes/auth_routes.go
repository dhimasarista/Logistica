package routes

// import (
// 	"log"
// 	"logistica/app/models"
// 	"time"

// 	"github.com/gofiber/fiber/v2"
// 	"github.com/gofiber/fiber/v2/middleware/session"
// )

// var store *session.Store = session.New()

// func AuthenticationRoutes(app *fiber.App) {
// 	// Middleware untuk session

// 	app.Get("/login", func(c *fiber.Ctx) error {

// 		return c.Render("login", fiber.Map{
// 			"Title": "LOGISTICA",
// 		})
// 	})

// 	app.Post("/login", func(c *fiber.Ctx) error {
// 		username := c.FormValue("username")
// 		password := c.FormValue("password")
// 		stayLoggedIn := c.FormValue("stay")
// 		var loggedIn bool = false

// 		users := models.User{}
// 		user := users.FindAll()[0]

// 		if user.Username != username {
// 			log.Println("Username Not Found")
// 			return c.Render("login", fiber.Map{
// 				"errors": []fiber.Map{
// 					{
// 						"message": "Username Not Found.",
// 					},
// 				},
// 			})
// 		} else if user.Password != password {
// 			log.Println("Password Incorrect")
// 			return c.Render("login", fiber.Map{
// 				"errors": []fiber.Map{
// 					{
// 						"message": "Password Incorrect.",
// 					},
// 				},
// 			})
// 		}
// 		session, err := store.Get(c)
// 		if err != nil {
// 			return c.Redirect("/505")
// 		}

// 		if stayLoggedIn == "true" {
// 			loggedIn = true
// 		}

// 		session.Set("Username", user.Username)
// 		session.Set("UserId", user.Password)
// 		session.Set("LoggedIn", loggedIn)
// 		session.SetExpiry(time.Minute * 60)
// 		store.CookieHTTPOnly = true
// 		store.CookieSecure = true
// 		session.Save()
// 		log.Println("Berhasil Loggin")
// 		return c.Redirect("/dashboard")
// 	})
// }

// func DeauthenticationRoutes(app *fiber.App) {
// 	app.Get("/logout", func(c *fiber.Ctx) error {
// 		session, _ := store.Get(c)
// 		session.Destroy()
// 		return c.Redirect("/login")
// 	})

// }

// // func IsAuthenticated(c *fiber.Ctx) error {
// // 	session, err := store.Get(c)
// // 	if err != nil {
// // 		return c.Redirect("/500")
// // 	}

// // 	username := session.Get("Username")
// // 	if username == "" {

// // 	}
// // }
