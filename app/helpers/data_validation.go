package helpers

import "github.com/gofiber/fiber/v2"

func IsEmpty(c *fiber.Ctx, data map[string]interface{}, requiredFields ...string) bool {
	for _, field := range requiredFields {
		if data[field] == "" {
			return true
		}
	}
	return false
}
