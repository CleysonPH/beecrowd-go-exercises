package main

import "fmt"

func main() {
	var n1, n2, n3, n4 float64
	fmt.Scanf("%f %f %f %f", &n1, &n2, &n3, &n4)

	media := ((n1 * 2) + (n2 * 3) + (n3 * 4) + (n4 * 1)) / float64(10)
	fmt.Printf("Media: %.1f\n", media)

	if media >= 7 {
		fmt.Println("Aluno aprovado.")
	} else if media < 5 {
		fmt.Println("Aluno reprovado.")
	} else {
		fmt.Println("Aluno em exame.")

		var exame float64
		fmt.Scanf("%f", &exame)
		fmt.Printf("Nota do exame: %.1f\n", exame)

		media = (media + exame) / 2
		if media >= 5 {
			fmt.Println("Aluno aprovado.")
		} else {
			fmt.Println("Aluno reprovado.")
		}
		fmt.Printf("Media final: %.1f\n", media)
	}
}
