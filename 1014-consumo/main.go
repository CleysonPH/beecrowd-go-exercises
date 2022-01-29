package main

import "fmt"

func main() {
	var distancia int
	var combustivelGasto float64
	fmt.Scanf("%d\n%f", &distancia, &combustivelGasto)

	consumoMedio := float64(distancia) / combustivelGasto

	fmt.Printf("%.3f km/l\n", consumoMedio)
}
