package files

import (
	"bufio"
	"fmt"
	"os"

	"github.com/MatiLang1/Godesde0/ejercicios"
)

var filename string = "./files/txt/tabla.txt"

//pasamos la ruta completa porque ejecutamos desde main.go, almacenamos la ruta en una variable

func GrabaTabla() { //va a crear un archivo con una tabla de multplicar
	var texto string = ejercicios.TabladeMultiplicar()
	archivo, err := os.Create(filename)
	//ésta funcion del SO llamada "create" crea un archivo, si ya existe lo borra y lo crea nuevamente de 0
	if err != nil {
		fmt.Println("Error al crear el archivo" + err.Error())
		return
		//si hay un error imprime que hay un error y con el return vacío salimos del método "Grabatabla"
	}
	fmt.Fprintln(archivo, texto) //graba el contenido de texto en archivo
	archivo.Close()
	//se debe cerrar el archivo (ya sea que fue abierto para grabar o leer)

}

func SumaTabla() { //va a agarrar cualquier archivo que le indiquemos y va a agregarle una tabla de multiplicar
	var texto string = ejercicios.TabladeMultiplicar()
	// if Append(filename, texto) == false {
	if !Append(filename, texto) { //mejor forma de hacer el if
		fmt.Println("Error al concatenar contenido")
	}
}

func Append(filen string, text string) bool {
	arc, err := os.OpenFile(filen, os.O_WRONLY|os.O_APPEND, 0644) //se crea un buffer (arch)
	//La función "OpenFile" del SO abre el archivo de forma/tipo append para concatenarle mas contenido

	if err != nil {
		fmt.Println("Error al agregar infomarción " + err.Error())
		return false
	}

	_, err = arc.WriteString(text + "\n") //en el buffer/archivo "arc", con la funcion "WriteString" pasandole un string lo graba en el archivo (en este caso el string es "text")

	if err != nil {
		fmt.Println("Error durante la escritura " + err.Error())
		return false
	}

	arc.Close()
	return true
}

func LeoArchivo() {
	archivo, err := os.Open(filename) //la función "Open" solo hace lectura
	if err != nil {
		fmt.Println("Error al leer archivo" + err.Error())
		return //sale del método
	}

	scanner := bufio.NewScanner(archivo) //los datos no los recibe de teclado sino de "archivo"
	for scanner.Scan() {                 // el scanner.Scan devuelve un Booleano
		registro := scanner.Text() //la función Text nos devuelve cada linea del archivo
		fmt.Println("> " + registro)
	}
	archivo.Close() //nombre del archivo (en este caso se llama "archivo") .Close para cerrar el archivo
}
