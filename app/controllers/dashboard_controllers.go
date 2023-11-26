package controllers

import (
	"fmt"
	"logistica/app/models"

	"github.com/gofiber/fiber/v2"
)

func DashboardRender(c *fiber.Ctx) error {
	var path string = c.Path()
	var username string = c.Cookies("user")
	fmt.Println(c.Cookies("user"))

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
}
