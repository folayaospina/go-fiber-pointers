package logic

import (
	"github.com/folayaospina/go-fiber-pointers/internal"
)

/*
This func recives a value *string that point to the memory
position of the argument recived.
*/
func PruebaPointer(person *internal.Person) string {
	return "Esto es prueba pointer con struct person: " + person.Name
}

func PruebaPointersArray(personas *internal.Arreglo) {

}
