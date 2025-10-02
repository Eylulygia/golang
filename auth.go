package main

import (
    "context"
    "time"
    "github.com/gofiber/fiber/v2"
    "github.com/golang-jwt/jwt/v4"
    "golang.org/x/crypto/bcrypt"
)

func AuthRoutes(app *fiber.App) {
    // Login
    app.Post("/login", func(c *fiber.Ctx) error {
        var data map[string]string
        if err := c.BodyParser(&data); err != nil {
            return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
        }

        collection := client.Database("library").Collection("users")
        var user User
        err := collection.FindOne(context.TODO(), map[string]string{"email": data["email"]}).Decode(&user)
        if err != nil {
            return c.Status(401).JSON(fiber.Map{"error": "Invalid email or password"})
        }

        if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data["password"])); err != nil {
            return c.Status(401).JSON(fiber.Map{"error": "Invalid email or password"})
        }

        token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
            "email": user.Email,
            "exp":   time.Now().Add(time.Hour * 24).Unix(),
        })

        tokenString, err := token.SignedString(JwtSecret)
        if err != nil {
            return c.Status(500).JSON(fiber.Map{"error": "Could not generate token"})
        }

        return c.JSON(fiber.Map{"message": "Login successful", "token": tokenString})
    })

    // Me
    app.Get("/me", AuthMiddleware, func(c *fiber.Ctx) error {
        email := c.Locals("email").(string)
        collection := client.Database("library").Collection("users")
        var user User
        err := collection.FindOne(context.TODO(), map[string]string{"email": email}).Decode(&user)
        if err != nil {
            return c.Status(404).JSON(fiber.Map{"error": "User not found"})
        }
        return c.JSON(user)
    })
}
