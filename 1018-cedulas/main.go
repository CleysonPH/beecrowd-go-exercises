package main

import "fmt"

func main() {
	var total int
	fmt.Scanf("%d", &total)

	fmt.Println(total)

	quantidade100 := total / 100
	total -= quantidade100 * 100
	quantidade50 := total / 50
	total -= quantidade50 * 50
	quantidade20 := total / 20
	total -= quantidade20 * 20
	quantidade10 := total / 10
	total -= quantidade10 * 10
	quantidade5 := total / 5
	total -= quantidade5 * 5
	quantidade2 := total / 2
	total -= quantidade2 * 2
	quantidade1 := total

	fmt.Printf("%d nota(s) de R$ 100,00\n", quantidade100)
	fmt.Printf("%d nota(s) de R$ 50,00\n", quantidade50)
	fmt.Printf("%d nota(s) de R$ 20,00\n", quantidade20)
	fmt.Printf("%d nota(s) de R$ 10,00\n", quantidade10)
	fmt.Printf("%d nota(s) de R$ 5,00\n", quantidade5)
	fmt.Printf("%d nota(s) de R$ 2,00\n", quantidade2)
	fmt.Printf("%d nota(s) de R$ 1,00\n", quantidade1)
}
