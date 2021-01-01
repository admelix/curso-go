package main

import "fmt"

var Calculo func(int, int) int = func(num1 int, num2 int) int {
	return num1 + num2
}

func main() {

	fmt.Printf("sumo 5 + 7: %d \n", Calculo(5, 7))

	// resta
	Calculo = func(num1 int, num2 int) int {
		return num1 - num2
	}

	fmt.Printf("La resta de 4 - 2 es: %d\n", Calculo(4, 2))

	// multiplicacion
	Calculo = func(num1 int, num2 int) int {
		return num1 * num2
	}

	fmt.Printf("multiplico 2 * 5: %d \n", Calculo(2, 5))

	Operaciones()

	/* closure */
	tablaDel := 2
	MiTabla := Tabla(tablaDel)
	for i := 1; i < 11; i++ {
		fmt.Println(MiTabla())
	}

}

func Operaciones() {
	resultado := func() int {
		var a int = 23
		var b int = 13
		return a + b
	}
	fmt.Println(resultado())
}

// Closure, mantiene las variables alejadas de hacking. Es para un nivel de seguridad. Isolation Code!
func Tabla(valor int) func() int {
	// datos fuera de la funcion
	numero := valor
	secuencia := 0
	return func() int {
		secuencia++
		return numero * secuencia
	}
}
