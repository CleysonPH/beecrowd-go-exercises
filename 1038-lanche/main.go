package main

import "fmt"

func main() {
	var cod, qtd int
	var total float64
	fmt.Scanf("%d %d", &cod, &qtd)

	switch cod {
	case 1:
		total = float64(qtd) * 4.0
	case 2:
		total = float64(qtd) * 4.5
	case 3:
		total = float64(qtd) * 5
	case 4:
		total = float64(qtd) * 2
	case 5:
		total = float64(qtd) * 1.5
	}

	fmt.Printf("Total: R$ %.2f\n", total)
}
