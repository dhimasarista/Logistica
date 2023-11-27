package controllers

import "github.com/gofiber/fiber/v2"

func InventoryRender(c *fiber.Ctx) error {
	// var path string = c.Path()
	return c.Redirect("/500?error=Page+is+under+construction", fiber.StatusMovedPermanently)
	// return c.Render("inventory", fiber.Map{
	// 	"path": path,
	// })
}
