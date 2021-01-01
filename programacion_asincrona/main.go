package main

import (
	"fmt"
	"strings"
	"time"
)

/* El runtime de go, en este caso, nos pide que ingresemos una leta. De forma asincrona, empieza a iterar las letras
de mi nombre, pero, si yo presiono una letra, automaticamente cierra el programa aunque no se haya terminado la iteracion
de mi nombre. Por eso es importante fijarse bien en la forma en que hacemos las funciones asincronas que solo con
poner la palabra go antes de una funcion, ya la estoy haciendo asincrona. */

func main() {

	go miNombreLento("Jose")
	fmt.Println("Estoy aqui!")
	var x string
	fmt.Scanln(&x)
}

func miNombreLento(nombre string) {
	letras := strings.Split(nombre, "")
	for _, letra := range letras {
		time.Sleep(1000 * time.Millisecond)
		fmt.Println(letra)
	}
}
