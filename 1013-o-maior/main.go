package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c int
	fmt.Scanf("%d %d %d", &a, &b, &c)

	maiorAB := ((a + b) + int(math.Abs(float64(a-b)))) / 2
	maiorBC := ((b + c) + int(math.Abs(float64(b-c)))) / 2
	maior := ((maiorAB + maiorBC) + int(math.Abs(float64(maiorAB-maiorBC)))) / 2

	fmt.Printf("%d eh o maior\n", maior)
}
