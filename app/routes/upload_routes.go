package routes

import (
	"encoding/hex"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
)

func FileManagement(app *fiber.App) {
	app.Post("/upload/image", func(c *fiber.Ctx) error {
		form, err := c.MultipartForm() // Init Multipartform
		if err != nil {
			log.Println(err)
			return c.JSON(fiber.Map{
				"status": fiber.StatusBadRequest,
				"error":  err.Error(),
			})
		}

		// Mengambil files dengan key image dari map
		files := form.File["image"]
		for _, file := range files {
			err := c.SaveFile(file, "./app/uploads/images/"+file.Filename)
			if err != nil {
				log.Println(err)
				return c.JSON(fiber.Map{
					"status": fiber.StatusBadRequest,
					"error":  err.Error(),
				})
			}
		}

		return c.JSON(fiber.Map{
			"status":  c.Response().StatusCode(),
			"message": "Image uploaded successfully",
		})
	})
	app.Delete("/delete/image/:filename", func(c *fiber.Ctx) error {
		// Mengubah data yang dikirim dalam bentuk hexadecimal menjadi byte
		byte, err := hex.DecodeString(c.Params("filename"))
		if err != nil {
			log.Println(err)
			return c.JSON(fiber.Map{
				"status": fiber.StatusBadRequest,
				"error":  err.Error(),
			})
		}
		// Mengubah byte menjadi string
		var filename string = string(byte)
		// Menghapus file
		err = os.Remove("./app/uploads/images/" + filename)
		if err != nil {
			log.Println(err)
			return c.JSON(fiber.Map{
				"status": fiber.StatusBadRequest,
				"error":  err.Error(),
			})
		}

		// Mengirim response 200 ke client
		return c.JSON(fiber.Map{
			"status":  c.Response().StatusCode,
			"message": "Success delete image",
		})
	})
}
