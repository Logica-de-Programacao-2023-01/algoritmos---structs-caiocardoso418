package main

import "fmt"

type Triangulo struct {
	base   float64
	altura float64
}

func calcularAreaTriangulo(t Triangulo) float64 {
	return (t.base * t.altura) / 2
}

func main() {
	triangulo := Triangulo{
		base:   7.2,
		altura: 6.2,
	}

	area := calcularAreaTriangulo(triangulo)
	fmt.Printf("Área do triângulo: %.2f\n", area)
}
