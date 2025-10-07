package main

import (
	"log"
	"fiber/config"
	"os"
	"fiber/routes"
	"github.com/joho/godotenv"
	"github.com/gofiber/fiber/v2"
)

func main() {

	 err := godotenv.Load()
    if err != nil {
        log.Println("⚠️  No .env file found")
    }


	app := fiber.New(fiber.Config{
		Prefork:       false,
		CaseSensitive: true,
		StrictRouting: true,
		ServerHeader:  "Fiber",
		AppName:       "Library Management System",
	})


	config.ConnectDB()

	routes.BookRoutes(app)
	routes.UserRoutes(app)

	for _, r := range app.GetRoutes() {
	log.Println(r.Method, r.Path)
}


	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" 
	}

	log.Printf("🚀 Server is running on 0.0.0.0:%s\n", port)
	log.Fatal(app.Listen("0.0.0.0:" + port))

}


