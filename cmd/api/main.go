package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/hojabri/backend/handlers"
	"github.com/hojabri/backend/repository"
	"log"
)

func main() {
	
	// Connect to database
	_, err := repository.Connect()
	if err != nil {
		log.Fatal(err)
	}
	
	// Create a new Fiber instance
	app := fiber.New()
	
	// Assign a hello message for root path
	app.Get("/", hello)
	
	// Use logger
	app.Use(logger.New())
	
	// Group user related APIs
	userGroup := app.Group("/user")
	
	userGroup.Get("/", handlers.GetAllUsers)
	userGroup.Get("/:id", handlers.GetSingleUser)
	userGroup.Post("/", handlers.AddNewUser)
	userGroup.Put("/:id", handlers.UpdateUser)
	userGroup.Delete("/:id", handlers.DeleteUser)
	
	err = app.Listen(":3000")
	if err != nil {
		 log.Fatal(err)
	}
}

func hello(c *fiber.Ctx) error {
	return c.SendString("Hello, World 👋!")
}