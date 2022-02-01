package main

import "fmt"

func main() {
	var x, y float64
	fmt.Scanf("%f %f", &x, &y)

	q := ""
	if x == 0 && y == 0 {
		q = "Origem"
	} else if x == 0 {
		q = "Eixo Y"
	} else if y == 0 {
		q = "Eixo X"
	} else if x > 0 {
		if y > 0 {
			q = "Q1"
		} else {
			q = "Q4"
		}
	} else if x < 0 {
		if y > 0 {
			q = "Q2"
		} else {
			q = "Q3"
		}
	}

	fmt.Println(q)
}
