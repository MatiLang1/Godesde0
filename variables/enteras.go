package variables //package, nombre de carpeta que la contiene
import "fmt"

func MuestroEnteros() {
	var intcomun int     //DECLARACION de variable (dado que usamos var)
	intde32 := int32(10) //ASIGNACION de variable, la variable toma el tipo de dato de lo que le estamos asignando
	intde64 := int64(100)

	fmt.Println("intcomun = ", intcomun)
	fmt.Println("intde32 = ", intde32)
	fmt.Println("intde64 = ", intde64)
}

//las variables no se inicializan en cero sino que lo hacen en su valor mas bajo
