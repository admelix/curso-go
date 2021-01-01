package main

var estado bool

func main() {

	estado = true
	if estado == true {
		println("El estado es verdadero 😁")
	} else {
		println("el estado es falso 😱")
	}

	cambio()

}

func cambio() {

	switch numero := 1111; numero {

	case 1:
		println(1)

	case 2:
		println(2)

	case 3:
		println(3)

	case 4:
		println(4)

	case 5:
		println(5)

	default:
		println("El numero es mayor a 5")

	}

}
