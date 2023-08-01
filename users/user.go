package users

import (
	"fmt"
	"time"

	"github.com/drodmen/godesde0/modelos"
)

func AltaUsuario() {
	u := new(modelos.User)                   //creando objeto
	u.AddUser(10, "Pablo", time.Now(), true) //usamos la función de agregar usuario
	fmt.Println(u)

}
