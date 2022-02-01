package main

import "fmt"

func main() {
	var a, b, c int
	var o1, o2, o3 int
	fmt.Scanf("%d %d %d", &a, &b, &c)

	if a < b && a < c {
		o1 = a
		if b < c {
			o2 = b
			o3 = c
		} else {
			o2 = c
			o3 = b
		}
	} else if b < a && b < c {
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

	fmt.Printf("%d\n%d\n%d\n", o1, o2, o3)
	fmt.Printf("\n%d\n%d\n%d\n", a, b, c)
}
