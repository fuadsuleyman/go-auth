package conrtollers

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/fuadsuleyman/go-auth/database"
	"github.com/fuadsuleyman/go-auth/models"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

const SecretKey = "supersecretkey"

type tokenClaims struct {
	jwt.StandardClaims
	Usertype string
}

func Register(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	responseMessage := ""

	var existsUser models.User

	fmt.Println("existUser before query:", existsUser)
	database.DB.Where("username = ?", data["username"]).First(&existsUser)
	fmt.Println("existUser after query:", existsUser)
	fmt.Println("existUser after username:", existsUser.Username)
	fmt.Println("existUser after password:", existsUser.Password)
	fmt.Println("existUser after id:", existsUser.Id)
	fmt.Println("existUser after Usertype:", existsUser.Usertype)

	if existsUser.Id > 0{
		responseMessage = fmt.Sprintf("This username is alredy exits!")
		c.Status(fiber.StatusOK)
		return c.JSON(fiber.Map{
			"message": responseMessage,
		})	
	}

	password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)

	user := models.User{
		Username: data["username"],
		Usertype: data["usertype"],
		Password: password,
	}

	database.DB.Create(&user)


	responseMessage = fmt.Sprintf("New user successfully created! username: %v, usertype: %v", user.Username, user.Usertype)
	c.Status(fiber.StatusCreated)
	return c.JSON(fiber.Map{
		"message": responseMessage,
	})
	
}

func Login(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	var user models.User

	database.DB.Where("username = ?", data["username"]).First(&user)

	if user.Id == 0 {
		c.Status(fiber.StatusOK)
		return c.JSON(fiber.Map{
			"warning": "user not found",
		})
	}

	if err := bcrypt.CompareHashAndPassword(user.Password, []byte(data["password"])); err != nil {
		c.Status(fiber.StatusOK)
		return c.JSON(fiber.Map{
			"warning": "Incorrect password",
		})
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			Issuer:    user.Username,
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
		user.Usertype,
	},
	)

	token, err := claims.SignedString([]byte(SecretKey))

	if err != nil {
		c.Status(fiber.StatusOK)
		return c.JSON(fiber.Map{
			"warning": "could not login",
		})
	}

	return c.JSON(token)
}


func User(c *fiber.Ctx) error {
	token := c.Get("Authorization")
	return c.JSON(fiber.Map{
		"Token from header": token,
	})
}