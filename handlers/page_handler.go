package handlers

import "github.com/gofiber/fiber/v2"

func ServeIndexPage(c *fiber.Ctx) error  {
	return c.SendFile("./view/index.html")
}