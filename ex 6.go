package main

import (
	"fmt"
	"time"
)

type Funcionario struct {
	nome    string
	salario float64
	idade   int
}

func aumentarSalario(f *Funcionario, porcentagem float64) {
	f.salario += f.salario * (porcentagem / 100)
}

func diminuirSalario(f *Funcionario, porcentagem float64) {
	f.salario -= f.salario * (porcentagem / 100)
}

func calcularTempoServico(f Funcionario) int {
	idadeInicioTrabalho := 18
	anosTrabalho := time.Now().Year() - (f.idade + idadeInicioTrabalho)
	return anosTrabalho
}

func main() {
	funcionario := Funcionario{
		nome:    "João",
		salario: 5000,
		idade:   30,
	}

	fmt.Println(" inicial:", funcionario.salario)

	aumentarSalario(&funcionario, 10)
	fmt.Println(" após aumento:", funcionario.salario)

	diminuirSalario(&funcionario, 5)
	fmt.Println(" após diminuição:", funcionario.salario)

	tempoServico := calcularTempoServico(funcionario)
	fmt.Println("Tempo de serviço:", tempoServico, "anos")
}
