package main

import (
	"fmt"

	//VÍDEO 3.25
	f "github.com/drodmen/godesde0/goroutines"
	//VÍDEO 3.24
	//d "github.com/drodmen/godesde0/deferPanic"
	//VÍDEO 3.23
	//e "github.com/drodmen/godesde0/ejer_interfaces" // la e esun ALIAS
	//"github.com/drodmen/godesde0/modelos"
	//"github.com/drodmen/godesde0/users" //VÍDEO 3.22
	//"github.com/drodmen/godesde0/arreglos_slices/mapas" // VÍDEO 3.21
	//"github.com/drodmen/godesde0/mapas" // VÍDEO 3.21
	//"github.com/drodmen/godesde0/arreglos_slices"	// VÍDEO 3.20
	//"github.com/drodmen/godesde0/funciones" // VÍDEO 3.18
	//"github.com/drodmen/godesde0/files" // VÍDEO 3.17
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
	//files.LeoArchivo()

	//VIDEO 3.18
	//funciones.Calculos()
	//funciones.LlamarClosure()

	//VÍDEO 3.19
	//funciones.Exponencia(2)

	//VÍDEO 3.20
	//arreglos_slices.MuestroArreglos()
	//arreglos_slices.MuestroSlices()
	//arreglos_slices.Capacidad()

	//VÍDEO 3.21
	//mapas.MostrarMapas()

	//VÍDEO 3.22
	//users.AltaUsuario()

	//VÍDEO 3.23
	// Pedro := new(modelos.Hombre)
	// e.HumanosRespirando(Pedro)

	// Maria := new(modelos.Mujer)
	// e.HumanosRespirando(Maria)

	// VÍDEO 3.25
	//d.VemosDefer()
	//d.EjemploPanic() //vale para abortar el software un recover recuperar, tras un panic

	// VÍDEO 3.25
	// go f.MiNombreLento("Diego Rodríguez") // añadiendo go delante, se hace asíncronamente la función que se escriba

	// fmt.Println("Estoy aquí")
	// var x string
	// fmt.Scanln(&x)

	//VÍDEO 3.26
	canal1 := make(chan bool)

	go f.MiNombreLento("Diego Rodríguez", canal1) // añadiendo go delante, se hace asíncronamente la función que se escriba
	defer func() {
		<-canal1
	}()
	fmt.Println("Estoy aquí")

}
