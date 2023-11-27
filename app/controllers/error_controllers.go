package controllers

import "github.com/gofiber/fiber/v2"

func NotFoundHandler(c *fiber.Ctx) error {
	var path string = c.Path()

	return c.Render("error-404", fiber.Map{
		"path": path,
	})
}

func ISEHandler(c *fiber.Ctx) error {
	var path string = c.Path()

	return c.Render("error_page", fiber.Map{
		"path": path,
	})
}
