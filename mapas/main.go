// un map es un elemento que se puede serializar. Es la llave valor de java o cualquier otro lenguaje.

package main

import "fmt"

func main() {
	/*  aqui decimos que creamos un map con un indice string y de segundo parametro, sale un string
	el numero 5, indica la cantidad de espacio en memoria que se va a reservar para crear maps. Puede
	dejarse sin eso y agregar n cantida de maps, pero, el costo en memoria es mayor. Si yo lo declaro al inicio
	es una practica mucho mas sana. */
	paises := make(map[string]string, 5)
	fmt.Println(paises)

	paises["mexico"] = "D.F"
	paises["Chile"] = "Santiago de chile"

	campeonato := map[string]int{
		"Barcelona":    23,
		"colo colo":    32,
		"chivas":       43,
		"Boca Juniors": 30}

	// adicionar elementos al map
	campeonato["River Plate"] = 22
	// borrar un elemento
	delete(campeonato, "River Plate")
	fmt.Println(campeonato)

	for equipo, puntaje := range campeonato {
		fmt.Printf("El equipo %s, tiene un puntaje de %d\n", equipo, puntaje)
	}
	// si quiero saber la existencia de un valor dentro del map, el %t es para boolean. Ver ese tipo de valores.
	puntaje, existe := campeonato["Mineiro"]
	fmt.Printf("El puntaje capturado %d, y el equipo existe??? %t \n", puntaje, existe)

}
