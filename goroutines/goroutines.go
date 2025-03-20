package goroutines

import (
	"fmt"
	"strings"
	"time"
)

func MiNombreLentoo(nombre string) {
	letras := strings.Split(nombre, "") //split separa el texto por el caracter indicado (en este caso vacío pero podria ser una letra o cualquier cosa)
	for _, letra := range letras {
		time.Sleep(1000 * time.Millisecond)
		fmt.Println(letra)
	}
}
