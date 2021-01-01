package main

import (
	"bufio"
	"fmt"
	"os"
)

var numero1 int
var numero2 int
var suma int
var leyenda string

func main() {

	print("ingresa el primer numero: ")
	fmt.Scanf("%d", &numero1)

	print("Ingrese el segundo numero: ")
	fmt.Scanf("%d", &numero2)

	print("Descripcion: ")

	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		leyenda = scanner.Text()
	}

	suma = numero1 + numero2
	println(leyenda, suma)

}
