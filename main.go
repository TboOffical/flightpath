package main

import (
	"log"

	"github.com/glebarez/sqlite"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

var db *gorm.DB
var config AppConfig
var err error

func main() {
	//Start DB
	db, err = gorm.Open(sqlite.Open("./fp.db"), &gorm.Config{})
	if err != nil {
		log.Panicln(err.Error())
	}

	err = db.AutoMigrate(&Path{}, &AppConfig{}, &ApiKey{})
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Database migrated")

	//load app config

	config = AppConfig{}
	err = config.loadConfig()
	if err != nil {
		log.Fatal(err)
	}

	config.runConfigActions()

	if !config.EnableDBEventLogging {
		log.Println("Event logging disabled, you should probably enable this!")
	}

	//log event chatter
	go printChatter()

	//load the paths

	err = buildPaths()
	if err != nil {
		log.Fatal(err)
	}

	//simulate router
	for {
	}

	//For dev only: disable the router for faster path build testing
	return

	//create router
	app := fiber.New()
	app.Use(apiKeyMiddleware)

	paths := app.Group("/paths")
	paths.Get("/", getPathsHandler)
	paths.Post("/create", createPathHandler)

	log.Println("Flightpath server started on :3000")
	err = app.Listen(":3000")
	if err != nil {
		log.Fatal(err)
	}
}
