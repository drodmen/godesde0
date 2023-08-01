package goroutines

import (
	"fmt"
	"strings"
	"time"
)

// Para goroutines
// func MiNombreLento(nombre string) {
// 	letras := strings.Split(nombre, "") //cadena separada
// 	for _, letra := range letras {
// 		time.Sleep(1000 * time.Millisecond) //para que demore un segundo
// 		fmt.Println(letra)
// 	}
// }

// Para chanel
func MiNombreLento(nombre string, canal1 chan bool) {
	letras := strings.Split(nombre, "") //cadena separada
	for _, letra := range letras {
		time.Sleep(1000 * time.Millisecond) //para que demore un segundo
		fmt.Println(letra)
	}
	canal1 <- true //esto dispara un trigger
}
