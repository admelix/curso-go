package main

import "log"

func main() {

	/* defer, no se lanza una vez detectado el error. Deja que todo el codigo se ejecute
	y, luego de terminar la ejecucion, sale de la funcion y en este caso, libera la memoria cerrando
	el archivo que no existe */

	// archivo := "prueba.txt"
	// f, err := os.Open(archivo)

	// defer f.Close()
	// if err != nil {
	// 	fmt.Println("Error Abriendo el archivo")
	// 	os.Exit(1)
	// }

	/* El panic es un texto que se lanza despues de hacer una validacion, pero, no tenemos control de el. Solo
	lanza el mensaje que nosotros ponemos y luego deja de ejecutarse el programa. */

	ejemploPanic()

}

func ejemploPanic() {
	/* el defer, en este caso se hizo con una funcion anonima porque en caso contrario,
	solo ejecuta una accion. En cambio, si lo usamos con una funcion anonima, lo que podemos hacer es darle mas instrucciones*/
	defer func() {
		/* el recover, lo que hace es que al saltar un panic, lo captura y lo agrega a la variable a la que
		se le asigna el recover. Luego, log, crea un texto y el %v, es para insertar variant y lo que hace es mostrar el valor de
		recover. ademas, hace un os.Exit(1) */

		reco := recover()
		if reco != nil {
			log.Fatalf("Ocurrio un error que genero el primer panic \n %v", reco)
		}
	}()

	a := 1
	if a == 1 {
		panic("Se encontro el valor 1")
	}
}
