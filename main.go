package main

import (
	"log"
	"fiber/config"
	"os"
	"fiber/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {

	app := fiber.New(fiber.Config{
		Prefork:       true,
		CaseSensitive: true,
		StrictRouting: true,
		ServerHeader:  "Fiber",
		AppName:       "Library Management System",
	})


	config.ConnectDB()

	routes.BookRoutes(app)

	for _, r := range app.GetRoutes() {
	log.Println(r.Method, r.Path)
}


	port := os.Getenv("PORT")
if port == "" {
    port = "3000" // default for local dev
}
log.Fatal(app.Listen(":" + port))

}


