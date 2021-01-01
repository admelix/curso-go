package main

import "fmt"

func main() {

	fmt.Println(calculo(5, 46))
	fmt.Println(calculo(1, 1, 1, 1, 1, 1))
	fmt.Println(calculo(2, 2, 3))
	fmt.Println(calculo(4, 4, 4, 4, 4, 4, 4))
}

func uno(numero int) (z int) {
	z = numero * 2
	return
}

func dos(numero int) (int, bool) {

	if numero == 1 {
		return 5, false
	} else {

		return 10, true
	}

}

// parametros variables
func calculo(numero ...int) int {
	total := 0

	for item, num := range numero {
		total += num
		fmt.Printf("\n Item %d", item)
	}
	return total
}
