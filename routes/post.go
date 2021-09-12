package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/parvusvox/miniblog/controllers"
)

func PostRoutesClear(route fiber.Router) {
	route.Get("", controllers.GetPosts)
	route.Get("/:id", controllers.GetPost)
}

func PostRoutesProtected(route fiber.Router) {
	route.Post("", controllers.CreatePost)
	route.Delete("/:id", controllers.DeletePost)
}
