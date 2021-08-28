package main

import (
    "log"
    "strings"

    "github.com/gofiber/fiber/v2"
    "github.com/gofiber/fiber/v2/middleware/logger"
    "github.com/gofiber/template/django"

    "github.com/parvusvox/miniblog/config"
    "github.com/parvusvox/miniblog/controllers"
    "github.com/parvusvox/miniblog/routes"
    "github.com/joho/godotenv"
)

func setupRoutes(app *fiber.App){
    app.Get("/", controllers.GetIndex)

    api := app.Group("/api")
    routes.PostRoutes(api.Group("/posts"))
}

func setupUtilFuncs(engine *django.Engine){
    engine.AddFunc("replaceSpaces", func(name string) string {
        return strings.ReplaceAll(name, " ", "_")
    })
}

func main() {
    engine := django.New("./views", ".html")
    setupUtilFuncs(engine)

    // setup templating engine
    app := fiber.New(fiber.Config {
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

    err = app.Listen(":3000")
    if err != nil {
        panic(err)
    }
}
