package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/parvusvox/miniblog/controllers"
)

func PostRoutes(route fiber.Router) {
	route.Get("", controllers.GetPosts)
	route.Post("", controllers.CreatePost)
	route.Delete("/:id", controllers.DeletePost)
	route.Get("/:id", controllers.GetPost)
}
