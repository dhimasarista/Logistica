package controllers

import (
	"log"
	"logistica/app/models"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func DaftarSuratRender(c *fiber.Ctx) error {
	var path string = c.Path()

	surat := models.Surat{}
	var suratMasuk []map[string]interface{}

	data, err := surat.FindAll() // data adalah models.Surat[]
	if err != nil {
		log.Println(err)
		// return c.Status(fiber.StatusInternalServerError).SendString("Terjadi kesalahan: " + err.Error())
		return c.Render("error-500", fiber.Map{
			"message": err,
		})
	}

	for i := 0; i < len(data); i++ {
		if data[i].JenisSurat.String == "masuk" {
			mapData := make(map[string]interface{})
			mapData["tanggalSurat"] = data[i].TanggalSurat.Time.Format("02012006")
			mapData["judulSurat"] = strings.ToUpper(data[i].JudulSurat.String)
			mapData["penerima"] = strings.ToUpper(data[i].Penerima.String)
			mapData["pengirim"] = strings.ToUpper(data[i].Pengirim.String)
			mapData["perihal"] = strings.ToTitle(data[i].Perihal.String)
			mapData["jenisSurat"] = strings.ToUpper(data[i].JenisSurat.String)
			suratMasuk = append(suratMasuk, mapData)
		}
	}

	// Mengirimkan halaman HTML yang dihasilkan ke browser
	return c.Render("daftarsurat", fiber.Map{
		"path":        path,
		"surat-masuk": suratMasuk,
	})
}
