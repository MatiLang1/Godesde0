package modelos

import (
	"time"
)

type User struct { //estructura llamada "User"
	Id        int
	Name      string
	createdat time.Time
	status    bool
}

// creamos una funcion por fuera de la estructura que contenga lo de la estructura
// se coloca (usuario NombreDeLaEstructuraReferenciada), "Adduser" es el nombre de la nueva función
func (usuario *User) AddUser(id int, name string, createdAt time.Time, status bool) {
	usuario.Id = id //colocar this.NombreDeVariable = VariableDeLaEstructura
	usuario.Name = name
	usuario.createdat = createdAt //usuario es una variable puede ser otro nombre
	usuario.status = status
}

//podriamos crear diversas funciones que hagan referencia a la estructura "User", para que cuando creemos el objeto "user" éste venga con todas sus funciones
//el * apunta a un putero (posicion de memoria)
