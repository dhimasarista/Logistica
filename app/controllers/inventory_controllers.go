package controllers

import "github.com/gofiber/fiber/v2"

func InventoryRender(c *fiber.Ctx) error {
	var path string = c.Path()
	var user string = "dhimasarista"
	return c.Render("inventory_page", fiber.Map{
		"path": path,
		"user": user,
	})
}
