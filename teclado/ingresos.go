package teclado

import (
	"bufio" // todo lo que tiene que ver para lectura, tratamiento de archivos, etc.
	"fmt"
	"os" // sistema operativo
	"strconv"
)

// declaramos variables

var numero1 int
var numero2 int
var leyenda string
var err error

func IngresoNumeros() {
	scanner := bufio.NewScanner(os.Stdin) // variable que proviene del paquete bufio, que va a leer desde os. Stdin es entrada que es teclado stdout es pantalla que es salida

	fmt.Println("Ingrese número 1: ")
	//comprobación de error
	if scanner.Scan() { //si...ingresa algo...
		numero1, err = strconv.Atoi(scanner.Text())
		if err != nil { //si error es distinto de nil (que es nada, vacío, nulo)
			panic("El dato ingresado es incorrecto" + err.Error()) //se aborta la app con un panic
		}
	}

	fmt.Println("Ingrese número 2: ")
	//comprobación de error
	if scanner.Scan() { //si...ingresa algo...
		numero2, err = strconv.Atoi(scanner.Text())
		if err != nil { //si error es distinto de nil (que es nada, vacío, nulo)
			panic("El dato ingresado es incorrecto" + err.Error()) //se aborta la app con un panic
		}
	}

	fmt.Println("Ingrese leyenda: ")
	//comprobación de error
	if scanner.Scan() { //si...ingresa algo...
		leyenda = scanner.Text()
		if err != nil { //si error es distinto de nil (que es nada, vacío, nulo)
			panic("El dato ingresado es incorrecto" + err.Error()) //se aborta la app con un panic
		}
	}

	fmt.Println(leyenda, numero1*numero2)
}
