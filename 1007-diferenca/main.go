package main

import "fmt"

func main() {
	var a, b, c, d, diferenca int
	fmt.Scanf("%d\n%d\n%d\n%d", &a, &b, &c, &d)
	diferenca = a*b - c*d
	fmt.Printf("DIFERENCA = %d\n", diferenca)
}
