package config

import (
    "context"
    "fmt"
    "log"
    "os"
    "time"

    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
    "go.mongodb.org/mongo-driver/mongo/readpref"
)

type MongoInstance struct {
    Client *mongo.Client
    DB *mongo.Database
}

var MI MongoInstance

func ConnectDB() {
    ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
    client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MONGO_URI")))
    defer cancel()

    if err != nil {
        log.Fatal(err)
    }

    err = client.Ping(ctx, readpref.Primary())
    if err != nil {
        log.Fatal(err)
    }


    MI = MongoInstance {
        Client: client,
        DB: client.Database(os.Getenv("DATABASE_NAME")),
    }

    fmt.Println("Database connected!")
}






























