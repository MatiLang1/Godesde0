package main

import (
	"github.com/MatiLang1/Godesde0/webserver"
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

	// files.LeoArchivo()

	// funciones.LlamarClosure()

	// funciones.Exponencial(2)

	// arreglos_slices.Capacidad()

	// users.AltaUsuario()

	// Pedro := new(modelos.Hombre)
	// ejer_interfaces.HumanosRespirando(Pedro)

	// Maria := new(modelos.Mujer)
	// ejer_interfaces.HumanosRespirando(Maria)

	// defer_panic.EjemploPanic()

	// go goroutines.MiNombreLentoo("Matías Lang")
	// fmt.Println("Estoy")
	// var x string
	// fmt.Scanln(&x) //al mandar el valor que me pide por teclado se corta la ejecución

	// canal1 := make(chan bool) //creamos un chan (canal), luego el tipo (bool en este caso)
	// go goroutines.MiNombreLentoo("Matías Lang", canal1)
	// fmt.Println("Estoy")
	// <-canal1 //ahora canal1 es el que envía info (es un await espera a que el canal1 termine la ejecucion, en este caso canal1 no se asigna a ninguna variable)

	webserver.MiWebServer()

}

//se importa la carpeta del package que contiene el codigo requerido (funciones, propiedades , etc)
