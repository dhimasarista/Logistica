package controllers

import "github.com/gofiber/fiber/v2"

func InventoryRender(c *fiber.Ctx) error {
	var path string = c.Path()

	return c.Render("error-404", fiber.Map{
		"path": path,
	})
}
