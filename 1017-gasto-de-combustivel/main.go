package main

import "fmt"

func main() {
	const KM_LITROS int = 12
	var tempoViagem, velocidadeMedia int
	fmt.Scanf("%d\n%d", &tempoViagem, &velocidadeMedia)

	gastoTotal := float64(tempoViagem*velocidadeMedia) / float64(KM_LITROS)

	fmt.Printf("%.3f\n", gastoTotal)
}
