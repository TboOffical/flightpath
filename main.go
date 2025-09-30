package main

import (
	"log"
	"time"

	"github.com/glebarez/sqlite"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

var db *gorm.DB
var config AppConfig
var err error

type FPInfoStruct struct {
	Version string
}

var FPInfo FPInfoStruct = FPInfoStruct{
	Version: "0.0.1",
}

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

	//run the init methods for all the nodes
	registerPushbullet()
	registerEmailOutlet()
	registerTextModifier()
	registerTimeTrigger()

	//log event chatter
	go printChatter()

	//load the paths

	err = buildPaths()
	if err != nil {
		log.Fatal(err)
	}

	//simulate router
	//for {
	//}

	//For dev only: disable the router for faster path build testing
	//return

	//create router
	app := fiber.New()
	app.Use(apiKeyMiddleware)

	v1 := app.Group("/v1")

	paths := v1.Group("/paths")
	paths.Get("/", getPathsHandler)
	paths.Post("/create", createPathHandler)
	paths.Post("/update", updatePathHandler)

	server := v1.Group("server")
	server.Get("/node_info", getDocsHandler)
	server.Get("/server_info", func(ctx *fiber.Ctx) error {
		t := time.Now()
		return ctx.JSON(fiber.Map{
			"server_info":     FPInfo,
			"server_time":     t.String(),
			"db_event_saving": config.EnableDBEventLogging,
		})
	})

	log.Println("Flightpath server started on :3000")
	err = app.Listen(":3000")
	if err != nil {
		log.Fatal(err)
	}
}
