package deferPanic

import (
	"fmt"
	"log"
)

func VemosDefer() {
	fmt.Println("Este es un primer mensaje")
	defer fmt.Println("Este es un mensaje final") // permite indicar al sistema, cuáles son las funciones finales a ejecutar, por ejemplo, el hacer un close de la base de datos, se usa defer

	fmt.Println("Este es el segundo mensaje")

}

func EjemploPanic() {
	defer func() {
		reco := recover()
		if reco != nil {
			log.Fatalf("Ocurrió un error que generó Panic \n %v", reco)
		}
	}()
	a := 1
	if a == 1 {
		panic("Se encontró el valor 1")
	}
}
