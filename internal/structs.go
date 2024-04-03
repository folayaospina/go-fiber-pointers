package internal

type Person struct {
	Name string `json: name`
}

type Arreglo struct {
	personas []Person
}