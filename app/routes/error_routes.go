package routes

import (
	"fmt"
	"logistica/app/controllers"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

func ErrorRoutes(app *fiber.App, store *session.Store) {
	app.Get("/error", func(c *fiber.Ctx) error {
		var path string = c.Path()
		var username string = controllers.GetSessionUsername(c, store)

		return c.Render("error_page", fiber.Map{
			"path": path,
			"user": username,
		})
	})
	app.Get("/500", func(c *fiber.Ctx) error {
		return c.Redirect("/error?code=500&title=Internal+Server+Error&message=We+will+fix+it+as+soon+as+possible...")
	})
	app.Get("/404", func(c *fiber.Ctx) error {
		return c.Redirect("/error?code=404&title=Page+Not+Found&message=It+looks+like+you+found+a+glitch+in+the+matrix...")
	})
	app.Get("/418", func(c *fiber.Ctx) error {
		return c.Redirect("/error?code=418&title=I`am+a+Teapot&message=hahahahahahhahahahahahahhahaha...")
	})
}

func InternalServerError(c *fiber.Ctx, message string) error {
	var messageFormatted = strings.Replace(message, " ", "+", -1)
	var path = fmt.Sprintf("/error?code=500&title=Internal+Server+Error&message=%s", messageFormatted)
	return c.Redirect(path)
}
