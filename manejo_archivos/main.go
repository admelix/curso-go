package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	leoArchivo()
	leoArchivo2()
	graboArchivo()
	graboArchivo2()
}

/*  nil es el null de go. El ReadFile, lee todo el archivo y lo cierra, pero, no puedo manipular los datos.*/

func leoArchivo() {
	archivo, err := ioutil.ReadFile("./prueba.txt")
	if err != nil {
		fmt.Println("Hubo un error de lectura")
	} else {
		fmt.Println(string(archivo))
	}
}

/* Con la funcion os y open, podemos leer el archivo y manipularlo con bufio y su scaner */
func leoArchivo2() {

	archivo, err := os.Open("./prueba.txt")
	if err != nil {
		fmt.Println("Hubo un error de lectura")
	} else {
		scanner := bufio.NewScanner(archivo)

		for scanner.Scan() {
			registro := scanner.Text()
			fmt.Printf("Linea > " + registro + "\n")
		}
	}
	archivo.Close()
}

func graboArchivo() {
	archivo, err := os.Create("./prueba2.txt")
	if err != nil {
		fmt.Println("Hubo un error de lectura")
		return
	}
	fmt.Fprintln(archivo, "La primera linea de un nuevo archivo creado ")
	fmt.Fprintln(archivo, "Segunda linea de un archivo creado.")
	archivo.Close()
}

func graboArchivo2() {
	fileName := "./prueba2.txt"
	if Append(fileName, "\n Esta es una tercera linea, agregada con el append. ") == false {
		fmt.Println("Error al agregar una linea")
	}
}

func Append(archivo string, texto string) bool {

	arch, err := os.OpenFile(archivo, os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Hubo un error en Append")
		return false
	}

	_, err = arch.WriteString(texto)
	if err != nil {
		fmt.Println("Hubo un error en Append")
		return false
	}
	return true
}
