package main

import (
	"fmt"
)

type Viagem struct {
	origem  string
	destino string
	data    string
	preco   float64
}

func main() {
	viagens := []Viagem{
		{origem: "São Paulo", destino: "Rio de Janeiro", data: "2023-06-01", preco: 200.0},
		{origem: "São Paulo", destino: "Belo Horizonte", data: "2023-06-15", preco: 300.0},
		{origem: "São Paulo", destino: "Curitiba", data: "2023-06-30", preco: 150.0},
	}

	viagemMaisCara := encontrarViagemMaisCara(viagens)
	fmt.Println("Viagem mais cara:")
	fmt.Println("Origem:", viagemMaisCara.origem)
	fmt.Println("Destino:", viagemMaisCara.destino)
	fmt.Println("Data:", viagemMaisCara.data)
	fmt.Println("Preço:", viagemMaisCara.preco)
}

func encontrarViagemMaisCara(viagens []Viagem) Viagem {
	viagemMaisCara := viagens[0]
	for _, viagem := range viagens {
		if viagem.preco > viagemMaisCara.preco {
			viagemMaisCara = viagem
		}
	}
	return viagemMaisCara
}
