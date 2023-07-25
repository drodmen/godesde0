package funciones

import "fmt"

func Exponencia(valor int) {
	if valor > 10000000 {
		return
	}

	fmt.Println(valor)
	Exponencia(valor * 2)
}
