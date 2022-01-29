package main

import "fmt"

func main() {
	const DIAS_ANO = 365
	const DIAS_MES = 30

	var idadeEmDias int
	fmt.Scanf("%d", &idadeEmDias)

	anos := idadeEmDias / DIAS_ANO
	idadeEmDias -= anos * DIAS_ANO
	meses := idadeEmDias / DIAS_MES
	idadeEmDias -= meses * DIAS_MES
	dias := idadeEmDias

	fmt.Printf("%d ano(s)\n", anos)
	fmt.Printf("%d mes(es)\n", meses)
	fmt.Printf("%d dia(s)\n", dias)
}
