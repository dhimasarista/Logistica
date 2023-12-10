package controllers

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
)

func UploadImage(c *fiber.Ctx) error {
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
}

func DeleteImage(c *fiber.Ctx) error {
	filename := c.Params("filename")

	// Menghapus file
	err := os.Remove("./app/uploads/images/" + filename)
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
}
