package main

import (
	"crypto/sha512"
	"errors"
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type ApiKey struct {
	gorm.Model
	Nickname   string
	ApiKeyHash string
}

func (ApiKey) addKeyToDB(key string, nick string) error {
	if len(key) == 0 {
		return errors.New("API key is empty")
	}

	hasher := sha512.New()
	hasher.Write([]byte(key))

	apiKeyHash := fmt.Sprintf("%x", hasher.Sum(nil))

	db.Create(&ApiKey{ApiKeyHash: apiKeyHash, Nickname: nick})

	return nil
}

func apiKeyMiddleware(c *fiber.Ctx) error {
	headers := c.GetReqHeaders()

	if headers["Token"] == nil {
		c.Status(http.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"success": false,
			"message": "No Api Key",
		})
	}

	hasher := sha512.New()

	apiKey := headers["Token"][0]
	hasher.Write([]byte(apiKey))

	apiKeyHash := fmt.Sprintf("%x", hasher.Sum(nil))

	var key ApiKey
	db.First(&key, "api_key_hash = ?", apiKeyHash)

	if key.ApiKeyHash != apiKeyHash {
		ip := c.IP()
		e(fmt.Sprint("An unauthorized API key was used by a device with ip", ip))
		c.Status(http.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"success": false,
			"message": "Api key not found",
		})
	}

	return c.Next()
}
