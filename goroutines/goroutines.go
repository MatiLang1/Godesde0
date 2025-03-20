package goroutines

import (
	"fmt"
	"strings"
	"time"
)

func MiNombreLentoo(nombre string, canal1 chan bool) {
	letras := strings.Split(nombre, "") //split separa el texto por el caracter indicado (en este caso vacío pero podria ser una letra o cualquier cosa)
	for _, letra := range letras {
		time.Sleep(1000 * time.Millisecond)
		fmt.Println(letra)
	}
	canal1 <- true //asignando el valor "true" al canal1 (se agina con <-)
}
