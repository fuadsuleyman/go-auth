package main

import (
	// "fmt"
	"fmt"

	"github.com/fuadsuleyman/go-auth/database"
	"github.com/fuadsuleyman/go-auth/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {

	if err := initConfig(); err != nil {
		logrus.Fatalf("error initializing configs: %s", err.Error())
	}
	  	
	database.Connect()

    app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	app.Use(
		logger.New(), // add Logger middleware
	  )


	routes.Setup(app)

    app.Listen(fmt.Sprintf(":%s", viper.GetString("port")))
}


func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}