package handlers

import (
	"go_fiber_restfull/config"
	"go_fiber_restfull/models"
	"go_fiber_restfull/validator"
	"time"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *fiber.Ctx) error {
	user := new(models.User)

	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error":"Cannot parse JSON"})
	}

	// hash password 
	hashPassword , err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil{
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error":"Cannot hash password"})
	}

	user.Password = string(hashPassword)

	// save to db
	saveUser := config.DB.Create(user)
	
	if saveUser.Error != nil{
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error":"Cannot create user"})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"Message":"Success create user"})
}

func Login(c *fiber.Ctx)  error {
	loginRequest := struct {
		Email string `json:"email"`
		Password string `json:"password"`
	}{}

	if err := c.BodyParser(&loginRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}


	var user models.User
	result := config.DB.Where("email = ?", loginRequest.Email).First(&user)
	if result.Error != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid Credentials"})
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginRequest.Password))
    if err != nil {
        return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid credentials"})
    }

	token, err := validator.GenerateJWT(user.ID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error":"Could not generate token"})
	}	

	cookie := new(fiber.Cookie)
	cookie.Name = "jwt"
	cookie.Value = token
	cookie.Expires = time.Now().Add(24 * time.Hour)
	cookie.HTTPOnly = true
	cookie.Secure = false
	cookie.SameSite = "Lax"
	c.Cookie(cookie)

	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{"message":"Login Successfully"})

}