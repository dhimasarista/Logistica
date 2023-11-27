package controllers

import "github.com/gofiber/fiber/v2"

func InventoryRender(c *fiber.Ctx) error {
	// var path string = c.Path()
	return c.Redirect("/error?code=418&title=I`am+a+Teapot&message=hahahahahahhahahahahahahhahaha...")
	// return c.Render("inventory", fiber.Map{
	// 	"path": path,
	// })
}
