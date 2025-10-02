package main

import (
    "context"
    "github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

func BookRoutes(app *fiber.App) {
    // Borrow
    app.Post("/borrow", AuthMiddleware, func(c *fiber.Ctx) error {
        email := c.Locals("email").(string)
        var data map[string]string
        if err := c.BodyParser(&data); err != nil {
            return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
        }
        bookName := data["book"]

		if bookName == "" {
            return c.Status(400).JSON(fiber.Map{"error": "Book name is required"})
        }

        collection := client.Database("library").Collection("users")
        var user User
        err := collection.FindOne(context.TODO(), map[string]string{"email": email}).Decode(&user)
        if err != nil {
            return c.Status(404).JSON(fiber.Map{"error": "User not found"})
        }

        if len(user.Books) >= 2 {
            return c.Status(409).JSON(fiber.Map{"error": "Max 2 books allowed"})
        }

		for _, b := range user.Books {
            if b == bookName {
                return c.Status(409).JSON(fiber.Map{"error": "You already borrowed this book"})
            }
        }

        // Kitap başka kullanıcıda mı?
        var otherUser User
        err = collection.FindOne(context.TODO(), bson.M{"books": bookName}).Decode(&otherUser)
        if err == nil {
            // err == nil → kitap bulundu demek
            return c.Status(409).JSON(fiber.Map{"error": "Book already borrowed by another user"})
        }

        user.Books = append(user.Books, bookName)
        _, err = collection.UpdateOne(context.TODO(),
            map[string]string{"email": email},
            map[string]interface{}{"$set": map[string]interface{}{"books": user.Books}},
        )
        if err != nil {
            return c.Status(500).JSON(fiber.Map{"error": "Could not borrow book"})
        }

          return c.Status(201).JSON(fiber.Map{"message": "Book borrowed successfully", "books":   user.Books,})
    })

    // Return
    app.Post("/return", AuthMiddleware, func(c *fiber.Ctx) error {
        email := c.Locals("email").(string)
        var data map[string]string
        if err := c.BodyParser(&data); err != nil {
            return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
        }
        bookName := data["book"]

        collection := client.Database("library").Collection("users")
        var user User
        err := collection.FindOne(context.TODO(), map[string]string{"email": email}).Decode(&user)
        if err != nil {
            return c.Status(404).JSON(fiber.Map{"error": "User not found"})
        }

        var updatedBooks []string
        for _, b := range user.Books {
            if b != bookName {
                updatedBooks = append(updatedBooks, b)
            }
        }

        _, err = collection.UpdateOne(context.TODO(),
            map[string]string{"email": email},
            map[string]interface{}{"$set": map[string]interface{}{"books": updatedBooks}},
        )
        if err != nil {
            return c.Status(500).JSON(fiber.Map{"error": "Could not return book"})
        }

        return c.JSON(fiber.Map{"message": "Book returned", "books": updatedBooks})
    })
}
