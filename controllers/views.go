package controllers

import (
	"github.com/gofiber/fiber/v2"
)

func GetIndex(c *fiber.Ctx) error {
	return c.Render("index", fiber.Map{})
}

func GetDetails(c *fiber.Ctx) error {
	return c.Render("detail", fiber.Map{
		"id": c.Params("id"),
	})
}

func GetSubmit(c *fiber.Ctx) error {
	return c.Render("how_to_submit", fiber.Map{})
}
