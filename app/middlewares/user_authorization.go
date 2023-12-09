package middlewares

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

// Menangani autorisasi user
func UserAuthorization(store *session.Store) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Mengalami alamat path
		var path string = c.OriginalURL()

		// Mengambil data sesi yang tersimpan
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

		// Jika pengguna sudah login, lanjutkan ke handler berikutnya
		if username != nil {
			return c.Next()
		}

		// Pengecualian path yang bisa diakses tanpa login
		if path == "/home" {
			return c.Next()
		}
		if path == "/login" {
			return c.Next()
		}
		if path == "/check-session" {
			return c.Next()
		}

		// Jika pengguna belum login, arahkan ke halaman login
		return c.Redirect("/login")
	}
}
