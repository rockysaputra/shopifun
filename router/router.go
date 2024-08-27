package router

import (
	"shopifun/handler"
	"shopifun/middleware"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/logger"
)

func SetupRoutes(app *fiber.App) {
	// middleware

	// check api
	api := app.Group("/api", logger.New())
	api.Get("/", handler.Hello)

	// user
	user := app.Group("/user")
	//register
	user.Post("/register", handler.Register)
	user.Post("/login", handler.Login)

	user.Get("/login-google", handler.LoginGoogle)
	user.Get("login-google/callback", handler.GoogleCallback)

	app.Use(middleware.Authorization)
	//protected route
	user.Get("/detail", handler.DetailProfile)
}
