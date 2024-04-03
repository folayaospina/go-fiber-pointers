package main

import (
	"github.com/gofiber/fiber/v2"
)

func main()  {
	app := fiber.New(fiber.Config{
		Prefork:       true,
		CaseSensitive: true,
		StrictRouting: true,
		ServerHeader:  "Fiber",
		AppName: "Test App v1.0.1",
	})
	
	SetupRoutes(app)
	
	app.Listen(":3000")
}