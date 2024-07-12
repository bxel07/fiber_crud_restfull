package routes

import (
	"go_fiber_restfull/handlers"

	"github.com/gofiber/fiber/v2"
)

func SetupRouter(app *fiber.App)  {
	api := app.Group("/api")
	v1 := api.Group("v1")

	v1.Get("/posts", handlers.GetPosts)
	v1.Get("/posts/:id", handlers.GetPost)
	v1.Post("/posts", handlers.CreatePost)
	v1.Put("/posts/:id", handlers.UpdatePost)
	v1.Put("/posts/:id", handlers.DeletePost)

	// page
	app.Get("/", handlers.ServeIndexPage)
}