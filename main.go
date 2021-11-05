package main

import (
	// "fmt"
	"github.com/fuadsuleyman/go-auth/database"
	"github.com/fuadsuleyman/go-auth/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	  	
	database.Connect()

    app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	app.Use(
		logger.New(), // add Logger middleware
	  )


	routes.Setup(app)

    app.Listen(":8000")
}