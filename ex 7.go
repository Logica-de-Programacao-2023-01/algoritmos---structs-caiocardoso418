package main

import "fmt"

type Animal struct {
	nome    string
	especie string
	idade   int
	som     string
}

func modificarSom(a *Animal, novoSom string) {
	a.som = novoSom
}

func imprimirAnimal(a Animal) {
	fmt.Printf("Nome: %s\n", a.nome)
	fmt.Printf("Espécie: %s\n", a.especie)
	fmt.Printf("Idade: %d\n", a.idade)
	fmt.Printf("Som: %s\n", a.som)
}

func main() {
	animal := Animal{
		nome:    "cleiton",
		especie: "cachorro",
		idade:   8,
		som:     "latido",
	}

	imprimirAnimal(animal)

	modificarSom(&animal, "Miado")
	fmt.Println("Após mudar o som:")
	imprimirAnimal(animal)
}
