package main

import (
	"fmt"
)

func main() {
	const SEGUNDOS_HORA = 3600
	const SEGUNDOS_MINUTO = 60
	var tempoEmSegundos int
	fmt.Scanf("%d", &tempoEmSegundos)

	horas := tempoEmSegundos / SEGUNDOS_HORA
	tempoEmSegundos -= horas * SEGUNDOS_HORA
	minutos := tempoEmSegundos / SEGUNDOS_MINUTO
	tempoEmSegundos -= minutos * SEGUNDOS_MINUTO
	segundos := tempoEmSegundos

	fmt.Printf("%d:%d:%d\n", horas, minutos, segundos)
}
