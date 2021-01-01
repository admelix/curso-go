package main

import "fmt"

type SerVivo interface {
	estaVivo() bool
}

type humano interface {
	respirar()
	pensar()
	comer()
	sexo() string
	estaVivo() bool
}

type animal interface {
	respirar()
	comer()
	EsCarnivoro() bool
	estaVivo() bool
}

type vegetal interface {
	ClasificacionVegetal() string
}

// humano

type hombre struct {
	edad       int
	altura     float32
	peso       float32
	respirando bool
	pensando   bool
	comiendo   bool
	esHombre   bool
	vivo       bool
}

type mujer struct {
	hombre
}

func (h *hombre) respirar()      { h.respirando = true }
func (h *hombre) comer()         { h.comiendo = true }
func (h *hombre) pensar()        { h.pensando = true }
func (h *hombre) estaVivo() bool { return h.vivo }
func (h *hombre) sexo() string {
	if h.esHombre == true {
		return "Hombre"
	}
	return "Mujer"

}

func HumanosRespirando(hu humano) {

	hu.respirar()
	fmt.Printf("soy un/a %s, y ya estoy respirando \n", hu.sexo())

}

// genero animal

type perro struct {
	respirando bool
	comiendo   bool
	carnivoro  bool
	vivo       bool
}

func (p *perro) respirar()         { p.respirando = true }
func (p *perro) comer()            { p.comiendo = true }
func (p *perro) EsCarnivoro() bool { return p.carnivoro }
func (p *perro) estaVivo() bool    { return p.vivo }

func AnimalesRespirar(an animal) {
	an.respirar()
	fmt.Println("soy un animal y estoy respirando")
}

func AnimalesCarnivoros(an animal) int {
	if an.EsCarnivoro() == true {
		return 1
	}
	return 0

}

func estaVivo(ev SerVivo) bool {

	return ev.estaVivo()

}

func main() {

	// Jose := new(hombre)
	// Jose.esHombre = true
	// HumanosRespirando(Jose)

	// Anita := new(mujer)
	// HumanosRespirando(Anita)

	totalCarnivoros := 0
	Dogo := new(perro)
	Dogo.carnivoro = true
	Dogo.vivo = true
	AnimalesRespirar(Dogo)
	totalCarnivoros = +AnimalesCarnivoros(Dogo)

	fmt.Printf("Total carnivoros: %d\n", totalCarnivoros)
	fmt.Printf("estoy vivo?: %t", estaVivo(Dogo))

}
