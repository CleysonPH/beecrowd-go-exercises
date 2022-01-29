package main

import "fmt"

func main() {
	const DISTANCIA_POR_MINUTO float64 = 0.5
	var distancia int
	fmt.Scanf("%d", &distancia)

	tempoEmMinutos := float64(distancia) / DISTANCIA_POR_MINUTO

	fmt.Printf("%d minutos\n", int(tempoEmMinutos))
}
