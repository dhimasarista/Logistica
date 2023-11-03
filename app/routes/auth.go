package routes

import (
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

func AuthRoutes(app *fiber.App) {

	store := session.New()

	var username string = "admin"
	var password string = "vancouver"
	app.Get("/login", func(c *fiber.Ctx) error {

		return c.Render("login", fiber.Map{
			"Title": "LOGISTICA",
		})
	})

	app.Post("/login", func(c *fiber.Ctx) error {
		session, err := store.Get(c)
		if err != nil {
			log.Println(err)
		}

		usernameForm := c.FormValue("username")
		passwordForm := c.FormValue("password")

		log.Println(usernameForm, passwordForm)

		if username == usernameForm && password == passwordForm {
			session.Set("username", usernameForm)
			session.SetExpiry(time.Second * 36000)
			if err := session.Save(); err != nil {
				log.Println(err)
			}
			log.Println("Berhasil Login")
			return c.Redirect("/dashboard")
		} else if username != usernameForm || password != passwordForm {
			log.Println("Gagal Login")
			return c.Render("login", fiber.Map{
				"error":   true,
				"message": "Login Gagal. Silahkan Coba Lagi",
			})
		} else {
			log.Println("Error")
			return c.Render("login", fiber.Map{
				"error":   true,
				"message": "Terjadi kesalahan internal server.",
			})
		}
	})

	app.Get("/logout", func(c *fiber.Ctx) error {
		store.Delete("authenticated")

		return c.Redirect("login")
	})
}
