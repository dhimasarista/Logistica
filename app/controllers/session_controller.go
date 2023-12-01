package controllers

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

func GetSessionUsername(c *fiber.Ctx, store *session.Store) string {
	session, err := store.Get(c)
	if err != nil {
		log.Println(err)
		return ""
	}

	username := session.Get("username")
	defer session.Save()

	return username.(string)
}
