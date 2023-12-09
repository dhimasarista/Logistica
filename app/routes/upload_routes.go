package routes

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func UploadFileRoutes(app *fiber.App) {
	app.Post("/upload/image", func(c *fiber.Ctx) error {
		const maxFileSize = 200 << 10 // 200kb dalam bentuk byte
		// Ambil filed dari form client
		image, err := c.FormFile("image")
		if err != nil {
			fmt.Println(err)
			return c.JSON(fiber.Map{
				"status": fiber.StatusBadRequest,
				"error":  err.Error(),
			})
		}

		// Memeriksa ukuran file
		if image.Size > maxFileSize {
			return c.JSON(fiber.Map{
				"status": fiber.StatusBadRequest,
				"error":  "File size exceeds the limit",
			})
		}

		err = c.SaveFile(image, "./uploads/images"+image.Filename)
		if err != nil {
			fmt.Println(err)
			return c.JSON(fiber.Map{
				"status": fiber.StatusBadRequest,
				"error":  err.Error(),
			})
		}

		fmt.Println(image)

		return c.JSON(fiber.Map{
			"status":  fiber.StatusAccepted,
			"message": "Image uploaded successfully",
		})
	})
}
