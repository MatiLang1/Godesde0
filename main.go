package main

import (
	"github.com/MatiLang1/Godesde0/files"
)

func main() {
	// variables.MuestroEnteros() //se ejecuta la funcion "Muestoenteros que está en la carpeta "variables"

	// variables.RestoVariables() //se ejecuta funcion del package "variables"

	// estado, texto := variables.ConviertoaTexto(1588)
	// fmt.Println(estado)
	// fmt.Println(texto)

	// valornumerico, textorespuesta := ejercicios.Ejercicio1("114")
	// fmt.Println(textorespuesta)
	// fmt.Println(valornumerico)

	// teclado.IngresoNumeros() //ingresas 2 numeros y un texto y te devuelve el texto con los n° multiplicados

	// iteraciones.Iterar()

	// fmt.Println(ejercicios.TabladeMultiplicar())

	// files.GrabaTabla()
	//guarda/graba la tabla de multiplicar dentro del archivo tabla.txt (esa tabla fue almacenada en la variable "texto" y luego escrita en archivo)

	// files.SumaTabla() //concatena tablas en el archivo "tabla.txt"

	files.LeoArchivo()
}

//se importa la carpeta del package que contiene el codigo requerido (funciones, propiedades , etc)
