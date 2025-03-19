package users

import (
	"fmt"
	"time"

	"github.com/MatiLang1/Godesde0/modelos"
)

func AltaUsuario() {
	u := new(modelos.User) //se coloca que "u" es de tipo modelos.User
	//U es una definicion hasta que se utiliza new para convertirla en objeto
	u.AddUser(10, "Pablo", time.Now(), true) //nos trae todas las propiedas y funciones del modelo "User"
	fmt.Println(u)
}
