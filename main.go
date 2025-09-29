package main

import (
    "context"
    "fmt"
    "log"
    "os"
    "time"

    "github.com/gofiber/fiber/v2"
    "github.com/joho/godotenv"
    "go.mongodb.org/mongo-driver/bson/primitive"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
    //"golang.org/x/crypto/bcrypt"
)

//MongoDB bağlantısını global olarak saklayan bir pointer.
var client *mongo.Client

type User struct {
    ID       primitive.ObjectID `bson:"_id,omitempty" json:"id"`
    Email    string             `json:"email"`
    Password string             `json:"password"`
    Books    []string           `json:"books"`
}

func main() {
	//.env yüklüyoruz
	err := godotenv.Load() //hata mesajini yakala 
	if err != nil {
		log.Fatal("Error loading .env")
	}
	//databse bağlanma islemi
	uri := os.Getenv("MONGODB_URI")
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

	client, err = mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
    	log.Fatal(err)
    }
	app := fiber.New()

	fmt.Println("🚀 Server running on http://localhost:3000")
    log.Fatal(app.Listen(":3000"))
	


}