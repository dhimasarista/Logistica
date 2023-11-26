package controllers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func IndexRender(c *fiber.Ctx) error {
	fmt.Println(c.Path())
	// Mengirimkan halaman HTML yang dihasilkan ke browser
	return c.Render("index", fiber.Map{
		"title":   "Logistica",
		"message": "Aplikasi Manajemen Logistik & Supply Chain",
	})
}
