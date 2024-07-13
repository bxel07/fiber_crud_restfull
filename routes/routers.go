package routes

import (
	"go_fiber_restfull/handlers"
	"go_fiber_restfull/middleware"

	"github.com/gofiber/fiber/v2"
)

func SetupRouter(app *fiber.App)  {
	api := app.Group("/api")
	auth := app.Group("/auth")

	// middleware for blog api
	v1 := api.Group("v1",  middleware.AuthMiddleware)
	

	// blog api
	v1.Get("/posts", handlers.GetPosts)
	v1.Get("/posts/:id", handlers.GetPost)
	v1.Post("/posts", handlers.CreatePost)
	v1.Put("/posts/:id", handlers.UpdatePost)
	v1.Delete("/posts/:id", handlers.DeletePost)


	// auth api
	auth.Post("/register", handlers.Register)
	auth.Post("/login", handlers.Login)
}