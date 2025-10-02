package main

import (
    "context"
    "fmt"
    "log"
    "os"
    "time"

    "github.com/gofiber/fiber/v2"
    "github.com/joho/godotenv"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

func main() {
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env")
    }

    //mongo bağlantısı, 10 saniye içinde bağlantı kurulmazsa iptal et, hata dön.
    uri := os.Getenv("MONGODB_URI")
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    client, err = mongo.Connect(ctx, options.Client().ApplyURI(uri))
    if err != nil {
        log.Fatal(err)
    }
    // gerçekten bağlandı mı test ediyoruz.
    err = client.Ping(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }
    app := fiber.New()

    // Route kayıtları
    RegisterRoute(app)
    AuthRoutes(app)
    DeleteRoute(app)
    BookRoutes(app)

    fmt.Println("Server running on http://localhost:3000")
    log.Fatal(app.Listen(":3000"))
}
