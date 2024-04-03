package main

import (
	"github.com/folayaospina/go-fiber-pointers/internal/handlers"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/home", handlers.HomePrincipalHandler)
	app.Post("/pointers", handlers.FuncionInicialDeApuntadoresEjemploHandler)
	app.Post("/array", handlers.ApuntadoresConArreglosHandler)
}