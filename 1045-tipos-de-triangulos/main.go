package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c float64
	var o1, o2, o3 float64
	fmt.Scanf("%f %f %f", &a, &b, &c)

	if a <= b && a <= c {
		o1 = a
		if b < c {
			o2 = b
			o3 = c
		} else {
			o2 = c
			o3 = b
		}
	} else if b <= a && b <= c {
		o1 = b
		if a < c {
			o2 = a
			o3 = c
		} else {
			o2 = c
			o3 = a
		}
	} else {
		o1 = c
		if a < b {
			o2 = a
			o3 = b
		} else {
			o2 = b
			o3 = a
		}
	}
	o1, o3 = o3, o1

	naoFormaTriangulo := o1 >= o2+o3
	if naoFormaTriangulo {
		fmt.Println("NAO FORMA TRIANGULO")
	} else {
		trianguloRetangulo := math.Pow(o1, 2) == math.Pow(o2, 2)+math.Pow(o3, 2)
		if trianguloRetangulo {
			fmt.Println("TRIANGULO RETANGULO")
		}

		trianguloObtusangulo := math.Pow(o1, 2) > math.Pow(o2, 2)+math.Pow(o3, 2)
		if trianguloObtusangulo {
			fmt.Println("TRIANGULO OBTUSANGULO")
		}

		trianguloAcutangulo := math.Pow(o1, 2) < math.Pow(o2, 2)+math.Pow(o3, 2)
		if trianguloAcutangulo {
			fmt.Println("TRIANGULO ACUTANGULO")
		}

		trianguloEquilatero := o1 == o2 && o2 == o3
		trianguloEscaleno := o1 != o2 && o2 != o3 && o3 != o1
		if trianguloEquilatero {
			fmt.Println("TRIANGULO EQUILATERO")
		} else if !trianguloEquilatero && !trianguloEscaleno {
			fmt.Println("TRIANGULO ISOSCELES")
		}
	}

}
