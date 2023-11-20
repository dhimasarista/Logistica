package routes

import (
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

func AuthRoutes(app *fiber.App) {
	store := session.New()

	app.Get("/login", func(c *fiber.Ctx) error {

		return c.Render("login", fiber.Map{
			"Title": "LOGISTICA",
		})
	})

	app.Post("/login", func(c *fiber.Ctx) error {
		session, err := store.Get(c)
		if err != nil {
			log.Println(err)
			return c.SendStatus(fiber.StatusInternalServerError)
		}

		usernameForm := c.FormValue("username")
		passwordForm := c.FormValue("password")

		log.Println(usernameForm, passwordForm)

		if isAuthenticated(usernameForm, passwordForm) {
			session.Set("username", usernameForm)
			session.SetExpiry(time.Second * 36000)

			if err := session.Save(); err != nil {
				log.Println(err)
				return c.SendStatus(fiber.StatusInternalServerError)
			}

			log.Println("Berhasil Login")
			return c.Redirect("/dashboard")
		}

		log.Println("Gagal Login")
		return c.Render("login", fiber.Map{
			"error":   true,
			"message": "Login Gagal. Silahkan Coba Lagi",
		})
	})

	app.Get("/logout", func(c *fiber.Ctx) error {
		sess, err := store.Get(c)
		if err != nil {
			panic(err)
		}

		store.Delete("username")

		// Destroy session
		if err := sess.Destroy(); err != nil {
			log.Println(err)
		}

		log.Println("Session End.")

		return c.Redirect("login")
	})
}

func isAuthenticated(username, password string) bool {
	// Gantilah logika autentikasi sesuai dengan kebutuhan Anda.
	// Ini adalah contoh sederhana. Anda harus memeriksa username dan password dalam basis data, biasanya.
	return username == "admin" && password == "vancouver"
}
