package main

import "fmt"

func main() {
	var a, b, c, media float64
	fmt.Scanf("%f\n%f\n%f", &a, &b, &c)
	media = ((a * 2) + (b * 3) + (c * 5)) / 10
	fmt.Printf("MEDIA = %.1f\n", media)
}