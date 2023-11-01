package controllers

import (
	"github.com/gofiber/fiber/v2"
)

func DashboardRender(c *fiber.Ctx) error {

	var path string = c.Path()

	// Mengirimkan halaman HTML yang dihasilkan ke browser
	return c.Render("dashboard", fiber.Map{
		"path": path,
	})
}
