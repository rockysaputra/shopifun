package main

import (
	config "shopifun/Config"
	model "shopifun/Model"
	"shopifun/router"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/log"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/gofiber/fiber/v3/middleware/limiter"
)

func main() {

	config.ConnectDB()

	db := config.DB

	db.AutoMigrate(&model.User{})
	app := fiber.New()
	app.Use(cors.New())

	app.Use(limiter.New(limiter.Config{
		Max:        100,
		Expiration: 30 * time.Second,
	}))

	router.SetupRoutes(app)

	log.Fatal(app.Listen(":3000"))
}
