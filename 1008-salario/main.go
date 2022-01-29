package main

import "fmt"

func main() {
	var funcionarioNumero, horasTrabalhadas int
	var valorHora, salario float64
	fmt.Scanf("%d\n%d\n%f\n", &funcionarioNumero, &horasTrabalhadas, &valorHora)
	salario = valorHora * float64(horasTrabalhadas)
	fmt.Printf("NUMBER = %d\n", funcionarioNumero)
	fmt.Printf("SALARY = U$ %.2f\n", salario)
}
