package main

import "fmt"

type Endereco struct {
	rua    string
	numero int
	cidade string
	estado string
}

type Pessoa struct {
	nome     string
	idade    int
	endereco Endereco
}

func imprimirEnderecoCompleto(p Pessoa) {
	fmt.Printf("Endereço de %s:\n", p.nome)
	fmt.Printf("Rua: %s, Número: %d\n", p.endereco.rua, p.endereco.numero)
	fmt.Printf("Cidade: %s, Estado: %s\n", p.endereco.cidade, p.endereco.estado)
}

func main() {
	endereco := Endereco{
		rua:    "SHIN Qi",
		numero: 7,
		cidade: "Brasília",
		estado: "DF",
	}

	pessoa := Pessoa{
		nome:     "Caio",
		idade:    30,
		endereco: endereco,
	}

	imprimirEnderecoCompleto(pessoa)
}
