package handlers

import (
	"context"
	"time"
	"os"
	"log"

	"github.com/gofiber/fiber/v2"
	"fiber/models"
	"fiber/config"
	"fiber/utils"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"go.mongodb.org/mongo-driver/bson"
)

func RegisterUser(c *fiber.Ctx) error {
	user := new(models.User)
	
	if err := c.BodyParser(user); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}

	validationErrors := utils.ValidateStruct(user)
	if validationErrors != nil {
		return c.Status(400).JSON(fiber.Map{
			"error":   "Validation failed",
			"details": validationErrors,
		})
	}
	
	userCollection := config.DB.Collection("users")
	var existing models.User
	err := userCollection.FindOne(context.TODO(), bson.M{"email": user.Email}).Decode(&existing)
	if err == nil {
		return c.Status(400).JSON(fiber.Map{"error": "Email already registered"})
	}

	hashed, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	user.Password = string(hashed)

	_, err = userCollection.InsertOne(context.TODO(), user)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to create user"})
	}

	return c.Status(201).JSON(fiber.Map{"message": "User registered successfully"})
}

func LoginUser(c *fiber.Ctx) error {
	var input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.BodyParser(&input); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}
		userCollection := config.DB.Collection("users")


	var user models.User
	err := userCollection.FindOne(context.TODO(), bson.M{"email": input.Email}).Decode(&user)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid credentials"})
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid credentials"})
	}

	secret := os.Getenv("JWT_SECRET")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": user.ID.Hex(),
		"exp":    time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		log.Println("JWT error:", err)
		return c.Status(500).JSON(fiber.Map{"error": "Failed to generate token"})
	}

	return c.JSON(fiber.Map{
		"token": tokenString,
	})
}


