package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mamilladruva0007/User-with-DOB-and-Calculated-Age/internal/handler"
)

func Setup(app *fiber.App, userHandler *handler.UserHandler) {
	api := app.Group("/api")

	users := api.Group("/users")
	users.Post("/", userHandler.CreateUser)
	users.Get("/", userHandler.ListUsers)
	users.Get("/:id", userHandler.GetUser)
	users.Put("/:id", userHandler.UpdateUser)
	users.Delete("/:id", userHandler.DeleteUser)
}
