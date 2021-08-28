package controllers

import (
    "os"
    "fmt"
    "time"

    "github.com/gofiber/fiber/v2"

    "github.com/parvusvox/miniblog/config"
    "github.com/parvusvox/miniblog/models"

    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
    "go.mongodb.org/mongo-driver/mongo"
)

func GetPosts(c *fiber.Ctx) error {
    postCollection := config.MI.DB.Collection(os.Getenv("POST_COLLECTION"))
    query := bson.D {{}}
    cursor, err := postCollection.Find(c.Context(), query)

    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map {
            "success": false, 
            "message": "Something went wrong",
            "error": err.Error(),
        })
    }

    var posts []models.Post = make([]models.Post, 0)
    err = cursor.All(c.Context(), &posts)
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "success": false,
            "message": "Something went wrong",
            "error": err.Error(),
        })
    }

    return c.Status(fiber.StatusOK).JSON(fiber.Map{
        "success": true,
        "data": fiber.Map{
            "posts": posts,
        },
    })
}

func CreatePost(c *fiber.Ctx) error {
    postCollection := config.MI.DB.Collection(os.Getenv("POST_COLLECTION"))
    data := new(models.Post)

    err := c.BodyParser(&data)
    if err != nil {
        fmt.Println(err)
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map {
            "success": false,
            "message": "Cannot parse JSON",
        })
    }

    data.ID = nil
    data.CreatedAt = time.Now()
    data.UpdatedAt = time.Now()

    result , err := postCollection.InsertOne(c.Context(), data)
    if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map {
            "success": false,
            "message": "Cannot insert post",
            "error": err, 
        })
    }

    post := &models.Post{}
    query := bson.D {{Key: "_id", Value: result.InsertedID}}
    postCollection.FindOne(c.Context(), query).Decode(post)

    return c.Status(fiber.StatusCreated).JSON(fiber.Map{
        "success": true,
        "data": fiber.Map{
            "post":post,
        },
    })
}

func GetPost(c *fiber.Ctx) error {
    postCollection := config.MI.DB.Collection(os.Getenv("POST_COLLECTION"))
    paramId := c.Params("id")
    id, err := primitive.ObjectIDFromHex(paramId)

    if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map {
            "success": false,
            "message": "cannot parse ID",
        })
    }

    post := &models.Post {}
    query := bson.D {{Key: "_id", Value: id}}
    err = postCollection.FindOne(c.Context(), query).Decode(post)

    if err != nil {
        return c.Status(fiber.StatusNotFound).JSON(fiber.Map {
            "success": false,
            "message": "Post not found",
            "error": err,
        })
    }

    return c.Status(fiber.StatusOK).JSON(fiber.Map {
        "success": true,
        "data": fiber.Map {
            "post":post,
        },
    })
}

func DeletePost(c *fiber.Ctx) error {
    postCollection := config.MI.DB.Collection(os.Getenv("POST_COLLECTION"))
    paramId := c.Params("id")
    id, err := primitive.ObjectIDFromHex(paramId)

    if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map {
            "success": false,
            "message": "cannot parse ID",
        })
    }

    query := bson.D{{Key:"_id", Value: id}}
    err = postCollection.FindOneAndDelete(c.Context(), query).Err()

    if err != nil {
        if err == mongo.ErrNoDocuments {
            return c.Status(fiber.StatusNotFound).JSON(fiber.Map {
                "success": false,
                "message": "Post not found",
                "error": err,
            })
        }

        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map {
            "success": false,
            "message": "Cannot delete post",
            "error": err,
        })
    }

    return c.SendStatus(fiber.StatusNoContent)
}




















