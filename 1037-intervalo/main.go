package main

import "fmt"

func main() {
	var n float64
	fmt.Scanf("%f", &n)

	if n >= 0 && n <= 25 {
		fmt.Println("Intervalo [0,25]")
	} else if n > 25 && n <= 50 {
		fmt.Println("Intervalo (25,50]")
	} else if n > 50 && n <= 75 {
		fmt.Println("Intervalo (50,75]")
	} else if n > 75 && n <= 100 {
		fmt.Println("Intervalo (75,100]")
	} else {
		fmt.Println("Fora de intervalo")
	}
}
