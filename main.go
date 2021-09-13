package main

import (
	"github.com/KinitaL/golang-api-v2/pkg/controllers"
	"github.com/KinitaL/golang-api-v2/pkg/middleware"
	"github.com/gofiber/fiber/v2"
)

func main() {
	//create an app
	app := fiber.New()

	//routing
	routes := app.Group("/api")
	auth := app.Group("/auth")

	//routes to auth
	auth.Post("register", controllers.Register)
	auth.Post("login", controllers.Login)
	auth.Post("delete", controllers.DeleteUser)

	//routes to manage content
	api := routes.Group("/collection")
	api.Use(middleware.Auth)
	api.Get("/", controllers.Get)
	api.Post("/", controllers.Post)
	api.Post("/:id?", controllers.Put)
	api.Delete("/", controllers.Delete)

	//start server on port 3002
	app.Listen(":3002")
}
