package controllers

import "github.com/gofiber/fiber/v2"

func IndexRender(c *fiber.Ctx) error {
	var path string = c.Path()
	if path == "/" {
		return c.Status(fiber.StatusMovedPermanently).Redirect("/dashboard")
	}
	return c.Next()
	// Mengirimkan halaman HTML yang dihasilkan ke browser
	// return c.Render("index", fiber.Map{
	// 	"path": path,
	// })
}
