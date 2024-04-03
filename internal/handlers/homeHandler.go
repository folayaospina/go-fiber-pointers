package handlers

import "github.com/gofiber/fiber/v2"

func HomePrincipalHandler(ctx *fiber.Ctx) error {
	return ctx.SendString("Hola, estas recibiendo tú primer string en go con Fiber")
}