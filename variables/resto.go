package variables

import (
	"fmt"
	"strconv" // para convertir strings, integers... entre ellos
	"time"    // para trabajar con fechas
)

/* variables aquí definidas, podrán ser vistas por cualquier función dentro del mismo paquete e incluso paquetes externos, siempre y cuando
   el nombre de la variable empiece por una mayúscula.
*/

var Nombre string
var Estado bool
var Sueldo float32
var Fecha time.Time

func RestoVariables() { //función pública
	Nombre = "Pedro"
	Estado = true
	Sueldo = 1577.66
	Fecha = time.Now()

	fmt.Println(Nombre)
	fmt.Println(Estado)
	fmt.Println(Sueldo)
	fmt.Println(Fecha)
}

func ConviertoATexto(numero int) (bool, string) {
	texto := strconv.Itoa(numero)
	return true, texto
}
