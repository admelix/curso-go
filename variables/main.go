package main

import "fmt"

var numero int
var texto string
var status bool

func main() {

	numero2, numero3, numero4, texto := 2, 5, 6, "soy un texto"

	fmt.Println(numero2)
	fmt.Println(numero3)
	fmt.Println(numero4)
	fmt.Println(texto)

	mostrarStatus()
}

func mostrarStatus() {
	status = true
	println(status)
}

// cuando empiezo en mayusculas una funcion, es porque es publica y la puedo exportar. Cuando empieza en minusculas,
// no puedo exportarla y es privada. Lo mismo pasa con las variables.
