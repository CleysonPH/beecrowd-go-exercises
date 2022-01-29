package main

import "fmt"

func main() {
	var raio float64
	fmt.Scanf("%f", &raio)
	var pi float64 = 3.14159
	var area float64 = pi * (raio * raio)
	fmt.Printf("A=%.4f\n", area)
}
