package main

import "fmt"

var result int

func main() {

	/* los middleware o interceptores, permiten ejecutar instrucciones comunes a varias funciones que reciben y devuelven
	los mismos tipos de variables. */

	fmt.Println("inicio")

	result = operacionesMidd(sumar)(2, 3)
	fmt.Println(result)

	result = operacionesMidd(restar)(4, 2)
	fmt.Println(result)

	result = operacionesMidd(multiplicar)(2, 3)
	fmt.Println(result)

}

func sumar(a, b int) int {
	return a + b
}

func restar(a, b int) int {
	return a - b
}

func multiplicar(a, b int) int {
	return a * b
}

func operacionesMidd(f func(int, int) int) func(int, int) int {

	return func(a, b int) int {
		fmt.Println("inicio de operacion")
		return f(a, b)
	}

}
