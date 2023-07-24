package main

import (
	//"fmt"

	"github.com/drodmen/godesde0/files" // VÍDEO 3.17
	//"github.com/drodmen/godesde0/ejercicios" // VÍDEO 3.16
	//"github.com/drodmen/godesde0/iteraciones" // VÍDEO 3.15
	//"github.com/drodmen/godesde0/teclado" // VÍDEO 3.14
	//"github.com/drodmen/godesde0/ejercicios" // VÍDEO 3.13
	//"github.com/drodmen/godesde0/variables"  // VÍDEO 3.9
)

func main() {
	// variables.MuestroEnteros() //vídeo 9
	// variables.RestoVariables() //vídeo 3.10
	/*estado, texto := variables.ConviertoATexto(1588)
	fmt.Println(estado)
	fmt.Println(texto)*/ //vídeo 3.11

	/* //Ejercicio 01 del vídeo 3.13
	numero, mensaje := ejercicios.DosValores("50")
	fmt.Println(numero)
	fmt.Println(mensaje)
	*/

	//VÍDEO 3.14
	//teclado.IngresoNumeros()

	//bucle - VÍDEO 3.15
	/*for {
		fmt.Println("hola")
		break
	}*/

	//iteraciones.Iterar() // video 3.15

	// vídeo 3.16 Ejercicio 02
	//ejercicios.TablasMultiplicar()

	//vídeo 3.17 Manejo de archivos en go
	//fmt.Println(ejercicios.TablasMultiplicar())

	//files.GrabaTabla()
	//files.SumaTabla()
	files.LeoArchivo()
}
