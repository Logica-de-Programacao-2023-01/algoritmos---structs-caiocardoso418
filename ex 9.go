package main

import "fmt"

type Aluno struct {
	nome  string
	idade int
	notas []float64
}

func main() {
	aluno := Aluno{
		nome:  "João",
		idade: 20,
		notas: []float64{7.5, 8.0, 9.5},
	}

	adicionarNota(&aluno, 8.5)
	removerNota(&aluno, 7.5)

	media := calcularMediaNotas(aluno)
	imprimirAluno(aluno, media)
}

func adicionarNota(a *Aluno, nota float64) {
	a.notas = append(a.notas, nota)
}

func removerNota(a *Aluno, nota float64) {
	for i, n := range a.notas {
		if n == nota {
			a.notas = append(a.notas[:i], a.notas[i+1:]...)
			break
		}
	}
}

func calcularMediaNotas(a Aluno) float64 {
	soma := 0.0
	for _, nota := range a.notas {
		soma += nota
	}
	return soma / float64(len(a.notas))
}

func imprimirAluno(a Aluno, media float64) {
	fmt.Println("Nome:", a.nome)
	fmt.Println("Idade:", a.idade)
	fmt.Println("Notas:", a.notas)
	fmt.Printf("Média das notas: %.2f\n", media)
}
