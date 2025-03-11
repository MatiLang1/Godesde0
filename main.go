package main

import (
	"fmt"

	"github.com/MatiLang1/Godesde0/ejercicios"
	"github.com/MatiLang1/Godesde0/variables" //el module de go.mod / la carpeta donde está lo que queremos traer en el package
)

func main() {
	variables.MuestroEnteros() //se ejecuta la funcion "Muestoenteros que está en la carpeta "variables"

	variables.RestoVariables() //se ejecuta funcion del package "variables"

	estado, texto := variables.ConviertoaTexto(1588)
	fmt.Println(estado)
	fmt.Println(texto)

	valornumerico, textorespuesta := ejercicios.Ejercicio1("114")
	fmt.Println(textorespuesta)
	fmt.Println(valornumerico)

}

//se importa la carpeta del package que contiene el codigo requerido (funciones, propiedades , etc)
