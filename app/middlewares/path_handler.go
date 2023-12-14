package middlewares

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

func PathHandler(store *session.Store) fiber.Handler {
	return func(c *fiber.Ctx) error {
		path := c.OriginalURL() // mengambil path
		// Mendapatkan sesi
		session, err := store.Get(c)
		if err != nil {
			log.Println(err)
			return err
		}

		// Mendapatkan nilai "username" dari sesi
		username := session.Get("username")

		// Menyimpan sesi setelah mendapatkan nilai "username"
		if err := session.Save(); err != nil {
			log.Println(err)
			return err
		}

		// Jika pengguna sudah login dan mengakses path /login
		// Alihkan ke path / atau /dashboard
		if path == "/login" && username != nil {
			return c.Redirect("/dashboard")
		}

		// di linux kode ini bekerja
		// di windows kode ini 404
		if path == "/" {
			return c.Redirect("/home")
		}

		return c.Next()
	}
}
