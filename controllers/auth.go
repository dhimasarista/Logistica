package controllers

import "github.com/gofiber/fiber/v2"

func LoginRender(c *fiber.Ctx) error {

	// Mengirimkan halaman HTML yang dihasilkan ke browser
	return c.Render("login", fiber.Map{
		"Title": "SIPSURAT",
	})
}
