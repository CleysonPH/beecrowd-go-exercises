package main

import "fmt"

func main() {
	var a, b, c, d int
	fmt.Scanf("%d %d %d %d", &a, &b, &c, &d)

	isBMaiorQueC := b > c
	isDMaiorQueA := d > a
	isSomaCDMaiorQueSomaAB := c+d > a+b
	isCDPositivos := c > 0 && d > 0
	isAPar := a%2 == 0

	if isBMaiorQueC && isDMaiorQueA && isSomaCDMaiorQueSomaAB && isCDPositivos && isAPar {
		fmt.Println("Valores aceitos")
	} else {
		fmt.Println("Valores nao aceitos")
	}
}
