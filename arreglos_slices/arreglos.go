package arreglos_slices

import "fmt"

//creamos tabla y go deja la variable vacía

//var tabla []int //se define un slice, porque un slice no tiene una cantidad de dimensión de elementos
//var tabla [10]int //esto es un array

var tabla [10]int = [10]int{10, 0, 59} //esto es un array, donde se declaran valores directamente

var matriz [20][30]int //así se crea una matriz bidimensional

func MuestroArreglos() {
	tabla[7] = 33
	tabla[2] = 22

	tabla2 := [10]string{"Pablo", "Juan", "Rafa"} //otra forma de declarar una serie de valores en un array
	fmt.Println(tabla)
	fmt.Println(tabla2)

	//otra forma de introducir valores en un array
	for i := 0; i < len(tabla); i++ {
		tabla[i] = i
		fmt.Println(tabla)
	}

	matriz[7][24] = 15
	fmt.Println(matriz)
}
