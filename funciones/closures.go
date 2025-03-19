package funciones

import "fmt"

func Tabla(valor int) func() int {
	//función llamda "Tabla" que recibe un valor int y devuelve una funcion anonima la cual devuelve un int

	numero := valor //se le asgina el valor que recibimos como parametro en "Tabla"
	secuencia := 0  //se inicializa en 0
	return func() int {
		secuencia++
		return numero * secuencia
	}

}

func LlamarClosure() {
	//interactua con la funcion que devuelve una función (con "Tabla" en este caso)
	tabladel := 2
	MiTabla := Tabla(tabladel) //MiTabla recibe el closure
	for i := 1; i < 11; i++ {
		fmt.Println(MiTabla())
	}

}
