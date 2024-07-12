package handlers

import (
	"go_fiber_restfull/config"
	"go_fiber_restfull/models"

	"github.com/gofiber/fiber/v2"
)

func GetPosts(c *fiber.Ctx) error {
	var posts []models.Post
	config.DB.Find(&posts)
	return c.JSON(posts)
}

func GetPost(c *fiber.Ctx) error {
	id := c.Params("id")
	var post models.Post
	result := config.DB.First(&post, id)
	if result.Error != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Post not found"})
	}
	return c.JSON(post)
}

func CreatePost(c *fiber.Ctx) error {
	post := new(models.Post)
	if err := c.BodyParser(post); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}

	// sanitize
	if err := post.Validate(); err != nil{
		return c.Status(400).JSON(fiber.Map{"error" : "failed to create post" })
	}

	// try to create new table data
	result := config.DB.Create(&post)
	if result.Error != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to create post"})
	}
	return c.JSON(post)
}

func UpdatePost(c *fiber.Ctx) error {
	id := c.Params("id")
	post := new(models.Post)
	if err := c.BodyParser(post); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}

	// Validate and sanitize all inputs
	if err := post.Validate(); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	result := config.DB.First(&models.Post{}, id)
	if result.Error != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Post not found"})
	}
	config.DB.Model(&models.Post{}).Where("id = ?", id).Updates(post)
	return c.JSON(post)
}

func DeletePost(c *fiber.Ctx) error {
	id := c.Params("id")
	result := config.DB.Delete(&models.Post{}, id)
	if result.Error != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Post not found"})
	}
	return c.SendStatus(204)
}