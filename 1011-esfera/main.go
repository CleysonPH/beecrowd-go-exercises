package main

import (
	"fmt"
	"math"
)

func main() {
	var raio float64
	const PI float64 = 3.14159
	fmt.Scanf("%f", &raio)
	volume := (4.0 / 3.0) * PI * math.Pow(raio, 3.0)
	fmt.Printf("VOLUME = %.3f\n", volume)
}
