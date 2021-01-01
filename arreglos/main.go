package main

import (
	"fmt"
)

// variable global
// var tabla [10]int

// var matriz [5][7]int

var matriz [5][7]int

func main() {

	// tabla[0] = 1
	// tabla[5] = 15

	// variable valida solo en el scope.
	// tabla := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// for i := 0; i < len(tabla); i++ {

	// 	fmt.Println(tabla[i])
	// }

	// matriz[3][5] = 1
	// fmt.Println(matriz)

	matriz := []int{2, 5, 4}
	fmt.Println(matriz)

	variante2()
	variante3()
	variante4()

}

func variante2() {
	elementos := [5]int{1, 2, 3, 4, 5}
	/* este es un slice. Lo que hace, es cortar un array y con el [3:] me indica que de la posicion 3 hasta el final
	van a ir adoptando los datos que contenga ese array. Por ente, en este caso, seria un array/vector de 2 posiciones.*/
	porcion := elementos[3:]
	fmt.Println(porcion)
}

func variante3() {
	/* el comando make, lo que hace, es crear un slice de una manera diferente. Primero, se indica un vector y el tipo de datos
	luego, el segundo parametro, indica la cantidad de espacio de ese vector y el tercer parametro, es hasta donde puede crecer
	este vector. Por eso es un slice, porque go, reserva 20 espacios en memoria pero nosotros, inicialmente usaremos 5 y,
	en dado caso que se necesiten mas espacios, llegaremos a un maximo de 20. Por eso se toma como slice, pero de una
	manera, mas sofisticada.
	*/
	elementos := make([]int, 5, 20)
	fmt.Printf("largo %d, capacidad %d", len(elementos), cap(elementos))
}

func variante4() {

	/* En este caso, creamos un vector de 0 posiciones con una capacidad de 0. Pero, el comando append, lo que hace
	es ir agregando elementos al vector dependiendo de la iteracion, aumentando las posiciones y, la capacidad se va midiendo
	de forma binaria. Por ejemplo, en este cacso que es 129 el maximo de iteraciones, la capacidad es de 256. Si fuese menor,
	seria de 2 ^ n en base a la cantidad de espacios creados en la iteracion. eso si, el costo en proceso es mas alto
	y se deberia utilizar solo en casos en los que no se tenga claro en cuanto podria ir aumentando un vector. */

	nums := make([]int, 0, 0)
	for i := 0; i < 129; i++ {
		nums = append(nums, i)
	}
	fmt.Printf("\n el largo %d y capacidad %d", len(nums), cap(nums))
}
