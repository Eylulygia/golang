package main

import (
    "fmt"
    "os"
    "github.com/gofiber/fiber/v2"
    "github.com/golang-jwt/jwt/v4"
)

var JwtSecret = []byte(os.Getenv("JWT_SECRET")) //daha güvenli

func AuthMiddleware(c *fiber.Ctx) error {
    tokenString := c.Get("Authorization")
    if tokenString == "" {
        return c.Status(401).JSON(fiber.Map{"error": "Missing token"})
    }

    // "Bearer" kısmını geç
    if len(tokenString) > 7 && tokenString[:7] == "Bearer " {
        tokenString = tokenString[7:]
    }

    token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, fmt.Errorf("unexpected signing method")
        }
        return JwtSecret, nil
    })

    if err != nil || !token.Valid {
        return c.Status(401).JSON(fiber.Map{"error": "Invalid token"})
    }

    claims := token.Claims.(jwt.MapClaims)
    c.Locals("email", claims["email"])

    return c.Next()
}
