package user

import "time"

type Usuario struct {
	Id        int
	Nombre    string
	FechaAlta time.Time
	Status    bool
}

func (this *Usuario) AltaUsuario(id int, Nombre string, fechaalta time.Time, status bool) {
	this.Id = id
	this.Nombre = Nombre
	this.FechaAlta = fechaalta
	this.Status = status
}
