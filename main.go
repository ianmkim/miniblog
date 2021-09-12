package main

import (
	"log"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/template/django"

	"github.com/joho/godotenv"
	"github.com/parvusvox/miniblog/config"
	"github.com/parvusvox/miniblog/controllers"
	"github.com/parvusvox/miniblog/routes"

	jwtware "github.com/gofiber/jwt/v2"
)

func setupRoutes(app *fiber.App) {
	app.Get("/", controllers.GetIndex)
	app.Get("/d/:id", controllers.GetDetails)
	app.Get("/submit", controllers.GetSubmit)
	app.Get("/article_test", func(c *fiber.Ctx) error {
		return c.Render("pol_post", fiber.Map{})
	})

	app.Post("/login", controllers.Login)
	app.Post("/hpr", controllers.HPR)

	api := app.Group("/api")
	routes.PostRoutesClear(api.Group("/posts"))

	app.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte(os.Getenv("SECRET")),
	}))
	routes.PostRoutesProtected(api.Group("/posts"))
}

func setupUtilFuncs(engine *django.Engine) {
	engine.AddFunc("replaceSpaces", func(name string) string {
		return strings.ReplaceAll(name, " ", "_")
	})
}

func main() {
	engine := django.New("./views", ".html")
	setupUtilFuncs(engine)

	// setup templating engine
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Use(logger.New())

	// setup static routes
	app.Static("/static", "./static")
	app.Static("/favicon.ico", "./favicon.ico")

	// load database credentials
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	config.ConnectDB()

	setupRoutes(app)

	port := os.Getenv("PORT")
	if port == "" {
		port = ":3000"
	} else {
		port = ":" + port
	}

	err = app.Listen(port)
	if err != nil {
		panic(err)
	}
}
