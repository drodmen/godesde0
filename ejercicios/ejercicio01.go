package ejercicios

import (
	"strconv"
)

func DosValores(texto string) (int, string) {
	numero, _ := strconv.Atoi(texto) // con _ ignoramos el error
	if numero > 100 {
		return numero, "Es mayor a 100"
	} else {
		return numero, "Es menor a 100"
	}
}

// Si queremos hacer tratamiento de error
/*
	numero, err := strconv.Atoi(texto)
	if err != nil{
		return 0, "Hubo un error" + err.Error()
	}

*/
