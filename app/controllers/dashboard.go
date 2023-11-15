package controllers

import (
	"logistica/app/models"

	"github.com/gofiber/fiber/v2"
)

func DashboardRender(c *fiber.Ctx) error {
	var path string = c.Path()
	var surat *models.Surat = &models.Surat{}
	var suratMasuk int = surat.CountSurat("masuk")
	suratKeluar := surat.CountSurat("keluar")

	// store := session.New()

	// sess, err := store.Get(c)
	// if err != nil {
	// 	panic(err)
	// }

	// Mengirimkan halaman HTML yang dihasilkan ke browser
	return c.Render("dashboard", fiber.Map{
		"path":               path,
		"total_surat_masuk":  suratMasuk,
		"total_surat_keluar": suratKeluar,
	})
}
