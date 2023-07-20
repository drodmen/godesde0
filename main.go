package main

import (
	"fmt"

	"github.com/drodmen/godesde0/variables"
)

func main() {
	// variables.MuestroEnteros() //vídeo 9
	// variables.RestoVariables() //vídeo 3.10
	estado, texto := variables.ConviertoATexto(1588)
	fmt.Println(estado)
	fmt.Println(texto)

}
