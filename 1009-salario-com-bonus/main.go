package main

import "fmt"

func main() {
	var nome string
	var salario, vendas, bonus, total float64
	fmt.Scanf("%s\n%f\n%f", &nome, &salario, &vendas)
	bonus = vendas * 0.15
	total = salario + bonus
	fmt.Printf("TOTAL = R$ %.2f\n", total)
}
