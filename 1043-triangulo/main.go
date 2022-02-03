package main

import "fmt"

func main() {
	var a, b, c float64
	fmt.Scanf("%f %f %f", &a, &b, &c)

	formaTriangulo := (a+b) > c && (a+c) > b && (c+b) > a
	if formaTriangulo {
		perimetro := a + b + c
		fmt.Printf("Perimetro = %.1f\n", perimetro)
	} else {
		area := ((a + b) * c) / 2
		fmt.Printf("Area = %.1f\n", area)
	}
}
