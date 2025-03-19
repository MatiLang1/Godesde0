package mapas

import "fmt"

func MostrarMapas() {
	paises := make(map[string]string) //(map[tipo de clave]tipo de valor)
	fmt.Println(paises)

	paises["Mexico"] = "D.F."
	paises["Argentina"] = "Buenos Aires"
	//se le pasa al mapa creado en "paises", la clave y el valor ["clave"] = ["valor"]
	fmt.Println(paises)              //imprime map[Argentina:Buenos Aires Mexico:D.F.]
	fmt.Println(paises["Argentina"]) //imprime el valor de la clave usada

	campeonato := map[string]int{
		"Barcelona":       39,
		"Real Madrid":     38,
		"Manchester City": 37,
	}

	fmt.Println(campeonato)

	for equipo, puntaje := range campeonato {
		fmt.Printf("Equipo %s, tiene un puntaje de %d \n", equipo, puntaje)
		//%d es para número de base 10 (procesa Int), %s para String , %t para Bool

	}

	delete(campeonato, "Real Madrid") //(NombreDelMapa, "clave"), la clave del elemento que queremos borrar
	fmt.Println(campeonato)

	puntaje, existe := campeonato["Juventus"] //usando la clave obtener su valor, cuando consultamos un mapa devuelve el valor de la clave buscado y un booleano que indica si ese valor existe o no
	fmt.Printf("El puntaje es: %d, y el equipo existe = %t \n", puntaje, existe)

	puntaje2, existe2 := campeonato["Manchester City"]
	fmt.Printf("El puntaje es: %d, y el equipo existe = %t \n", puntaje2, existe2)

}
