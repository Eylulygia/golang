package main

import (
    "context"
    "github.com/gofiber/fiber/v2"
    "golang.org/x/crypto/bcrypt"
)

func RegisterRoute(app *fiber.App) {
    app.Post("/register", func(c *fiber.Ctx) error {
        var data map[string]string
        if err := c.BodyParser(&data); err != nil {
            return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
        }

        hashedPassword, err := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)
        if err != nil {
            return c.Status(500).JSON(fiber.Map{"error": "Password hashing failed"})
        }

        user := User{
            Email:    data["email"],
            Password: string(hashedPassword),
            Books:    []string{},
        }

        collection := client.Database("library").Collection("users")
        _, err = collection.InsertOne(context.TODO(), user)
        if err != nil {
            return c.Status(500).JSON(fiber.Map{"error": "Could not create user"})
        }

        return c.Status(201).JSON(fiber.Map{"message": "User registered successfully"})
    })
}
