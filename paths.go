package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"gorm.io/gorm"
)

type Path struct {
	gorm.Model
	Data string
}

type createPathParams struct {
	PathData string `json:"path_data"`
}

func getPathsHandler(c *fiber.Ctx) error {
	var paths []Path
	db.Find(&paths)

	return c.JSON(paths)
}

func createPathHandler(c *fiber.Ctx) error {
	p := createPathParams{}

	if err := c.BodyParser(&p); err != nil {
		return err
	}

	if len(p.PathData) == 0 {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	//create the path
	path := Path{
		Data: p.PathData,
	}

	tx := db.Create(&path)
	if tx.Error != nil {
		log.Errorf(tx.Error.Error())
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(fiber.Map{
		"success": true,
		"id":      path.ID,
	})
}
