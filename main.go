package main

import (
	"log"
	"fiber/config"
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


	log.Fatal(app.Listen(":3000"))
}


