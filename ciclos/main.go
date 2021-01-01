package main

import "fmt"

func main() {

	// i := 1
	// for i < 10 {
	// 	println(i)
	// 	i++
	// }

	// for i := 1; i <= 10; i++ {
	// 	println(i)
	// }

	// numero := 0

	// for {
	// 	println("Continuo")
	// 	println("Ingrese un numero secreto")
	// 	fmt.Scanf("%d", &numero)

	// 	if numero == 100 {
	// 		break
	// 	}
	// }

	// var i = 0
	// for i < 10 {
	// 	fmt.Printf("\n valor de i: %d", i)
	// 	if i == 5 {
	// 		fmt.Printf(" multiplicamos por 2 \n")
	// 		i *= 2
	// 		continue
	// 	}
	// 	fmt.Printf("  Paso por aqui \n")
	// 	i++
	// }

	var i int = 0

RUTINA:
	for i < 10 {
		if i == 4 {
			i += 2
			fmt.Printf(" Voy a rutina sumando 2 a i")
			goto RUTINA
		}
		fmt.Printf("\n valor de i: %d", i)
		i++
	}

}
