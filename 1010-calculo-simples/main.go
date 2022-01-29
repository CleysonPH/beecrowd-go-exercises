package main

import "fmt"

func main() {
	var cod1, cod2, num1, num2 int
	var valor1, valor2, total float64
	fmt.Scanf("%d %d %f", &cod1, &num1, &valor1)
	fmt.Scanf("%d %d %f", &cod2, &num2, &valor2)
	total = float64(num1)*valor1 + float64(num2)*valor2
	fmt.Printf("VALOR A PAGAR: R$ %.2f\n", total)
}
