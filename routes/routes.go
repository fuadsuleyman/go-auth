package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/fuadsuleyman/go-auth/conrtollers"
)


func Setup(app *fiber.App) {
	app.Post("/api/v1.0/register", conrtollers.Register)
}