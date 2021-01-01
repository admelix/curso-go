package main

import (
	"fmt"
	"time"

	us "./estructuras/usuarios"
)

// herencia de la estructura Usuario
type nuevoUsuario struct {
	us.Usuario
}

func main() {

	user := new(nuevoUsuario)
	user.AltaUsuario(1, "Jose Sakuda", time.Now(), true)

	fmt.Println(user.Usuario)

}
