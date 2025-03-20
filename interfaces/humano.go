package interfaces

type Humano interface { //creando interface "Humano"
	//en las interfaces no se colocan propiedades ni variables ni tipo de datos, solo se definen las funciones

	Respirar()
	Pensar()
	Comer()
	Sexo() string
}
