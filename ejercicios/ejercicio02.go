/*
	Crear una función pública que pida por teclado un nº, valide i hay error o no
	En base al nº recibido crear la tabla numérica del mismo. De 1 a 0 y la muestre por pantalla.
	En Main, llamar a dicha función
	Grabar, ejecutar y luego subir todo a Github
*/

package ejercicios

import (
	"bufio" // todo lo que tiene que ver para lectura, tratamiento de archivos, etc.
	"fmt"
	"os"      // sistema operativo
	"strconv" //convertir valores
)

var num1 int // número que será introducido por el usuario desde teclado
var err error

func TablasMultiplicar() {
	scanner := bufio.NewScanner(os.Stdin) // variable que proviene del paquete bufio, que va a leer desde os. Stdin es entrada que es teclado stdout es pantalla que es salida

	//Validación de la introducción del valor
	for { //esto es un bucle infinito, que se parará al entrar en el break
		fmt.Println("Ingrese número del que desea la tabla de multiplicar: ")
		if scanner.Scan() { //Si ingresa algo
			num1, err = strconv.Atoi(scanner.Text()) //lo convertimos de texto a número entero
			if err != nil {                          //si error es distinto de nil (que es nada, vacío, nulo)
				fmt.Println("El valor introducido no es válido, por favor, introduzca un número entero.")
				continue // se hace un continue
			} else {
				break // en caso contrario, se sale del bucle infinito
			}
		}
	}

	for i := 1; i <= 10; i++ {
		fmt.Println(num1, "x", i, " = ", num1*i)
	}
}
