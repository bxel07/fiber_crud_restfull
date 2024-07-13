package main

import(
	"go_fiber_restfull/config"
	"go_fiber_restfull/models"
	"go_fiber_restfull/routes"


	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main()  {

	
	app := fiber.New()


	// add CORS middleware
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,HEAD,PUT,DELETE,PATCH",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	// setup db and auto migrations
	config.ConnectDB()
	config.DB.AutoMigrate(&models.Post{}, &models.User{})


	// setup router
	routes.SetupRouter(app)
	app.Listen(":5000")
}