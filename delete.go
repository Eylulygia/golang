package main

import (
    "context"
    "github.com/gofiber/fiber/v2"
)

func DeleteRoute(app *fiber.App) {
    app.Delete("/delete", AuthMiddleware, func(c *fiber.Ctx) error {
        email := c.Locals("email").(string)
        collection := client.Database("library").Collection("users")
        res, err := collection.DeleteOne(context.TODO(), map[string]string{"email": email})
        if err != nil {
            return c.Status(500).JSON(fiber.Map{"error": "Could not delete user"})
        }
        if res.DeletedCount == 0 {
            return c.Status(404).JSON(fiber.Map{"error": "User not found"})
        }
        return c.JSON(fiber.Map{"message": "User deleted successfully"})
    })
}
