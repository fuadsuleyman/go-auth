package conrtollers

import (
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/fuadsuleyman/go-auth/database"
	"github.com/fuadsuleyman/go-auth/helper"
	"github.com/fuadsuleyman/go-auth/models"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"
)

// const SecretKey = "supersecretkey"

type tokenClaims struct {
	jwt.StandardClaims
	Usertype string
}


func Register(c *fiber.Ctx) error {
	var data map[string]string


	if err := c.BodyParser(&data); err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "message": err.Error(),
        })
    }

	var existsUser models.User

	// check if entered username exists in db or not
	database.DB.Where("username = ?", data["username"]).First(&existsUser)
	if existsUser.Id > 0{
		c.Status(fiber.StatusOK)
		return c.JSON(fiber.Map{
			"message": "This username is alredy exits!",
		})	
	}

	// hash password
	password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)

	user := models.User{
		Username: data["username"],
		Usertype: data["usertype"],
		Password: password,
	}

	errors := helper.ValidateStruct(user)
    if errors != nil {
		is_user := false
		is_type := false
		for _, val := range errors{
			if val.FailedField == "User.Username"{
				is_user = true
			} else if val.FailedField == "User.Usertype" {
				is_type = true
			}
		}
		if is_user && is_type{
			return c.JSON(fiber.Map{
				"warning": "username and usertype fields are required!",
			})
		} else if is_user {
			return c.JSON(fiber.Map{
				"warning": "username field is required!",
			})
		} else if is_type {
			return c.JSON(fiber.Map{
				"warning": "usertype field is required!",
			})
		}
       return c.JSON(errors)
    }

	// check usertype
	if data["usertype"] != "1" && data["usertype"] != "2" && data["usertype"] != "3"{
		return c.JSON(fiber.Map{
			"warning": "Wrong user type!",
		})
	}

	// check password
	if(data["password"] == ""){
		return c.JSON(fiber.Map{
			"warning": "password field is required!",
		})
	}

	// create user
	database.DB.Create(&user)

	responseMessage := fmt.Sprintf("New user successfully created! username: %v, usertype: %v", user.Username, user.Usertype)
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

	// check user and password
	warning := false
	if user.Id == 0 {
		warning = true
	}
	if err := bcrypt.CompareHashAndPassword(user.Password, []byte(data["password"])); err != nil {
		warning = true
	}
	if warning{
		c.Status(fiber.StatusOK)
		return c.JSON(fiber.Map{
			"warning": "username or password is incorrect!",
		})
	}

	// token claims data
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			Issuer:    user.Username,
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
		user.Usertype,
	},
	)

	token, err := claims.SignedString([]byte(viper.GetString("secret_ket")))

	if err != nil {
		c.Status(fiber.StatusOK)
		return c.JSON(fiber.Map{
			"warning": "could not login",
		})
	}
	return c.JSON(fiber.Map{
		"token": token,
	})
}


func User(c *fiber.Ctx) error {
	token := c.Get("Authorization")
	return c.JSON(fiber.Map{
		"Token from header": token,
	})
}