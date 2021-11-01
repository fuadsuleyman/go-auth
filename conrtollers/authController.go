package conrtollers

import (
	"github.com/fuadsuleyman/go-auth/models"
	"github.com/fuadsuleyman/go-auth/database"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)

	user := models.User{
		Username: data["username"],
		Usertype: data["usertype"],
		Password: password,
	}

	database.DB.Create(&user)

	return c.JSON(user)
}
