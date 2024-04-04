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
	personas := new(internal.Arreglo)

	if err := ctx.BodyParser(&personas); err != nil {
		return err
	}

	for _, v := range personas {
		p
	}

	return nil
}

func New(arreglo internal.Arreglo) {
	panic("unimplemented")
}
