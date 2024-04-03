package handlers

import (
	"github.com/folayaospina/go-fiber-pointers/internal"
	"github.com/folayaospina/go-fiber-pointers/internal/logic"
	"github.com/gofiber/fiber/v2"
)

func FuncionInicialDeApuntadoresEjemploHandler(ctx *fiber.Ctx) error {
	var person internal.Person

	if err := ctx.BodyParser(&person); err != nil {
		return err
	}

	err := ctx.SendString(logic.PruebaPointer(&person))
	if err != nil {
		return fiber.ErrInternalServerError // Manejo de error, dependiendo de la lógica de tu aplicación

	}

	return nil
}

func ApuntadoresConArreglosHandler(ctx *fiber.Ctx) error {
	var personas internal.Arreglo
/* 
	if err := ctx.BodyParser(&personas); err ! */= nil {
		return err
	}
	
	body := ctx.Body()
	reader := bytes.NewReader(body)
	lectura, err := io.ReadAll(reader)

	if err != nil {
		return err
	}

	return nil
}
