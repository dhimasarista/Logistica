package routes

import (
	"fmt"
	"logistica/app/models"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

func DashboardRoutes(app *fiber.App, store *session.Store) {
	app.Get("/dashboard", func(c *fiber.Ctx) error {
		var path string = c.Path()
		session, _ := store.Get(c)
		username := session.Get("username")
		defer session.Save()
		fmt.Println("/dashboard:", username)

		var surat *models.Surat = &models.Surat{}
		var suratMasuk int = surat.CountSurat("masuk")
		suratKeluar := surat.CountSurat("keluar")

		// Mengirimkan halaman HTML yang dihasilkan ke browser
		return c.Render("dashboard", fiber.Map{
			"path":               path,
			"user":               username,
			"total_surat_masuk":  suratMasuk,
			"total_surat_keluar": suratKeluar,
		})
	},
	)
}
