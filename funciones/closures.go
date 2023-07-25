package funciones

import (
	"fmt"
)

func tabla(valor int) func() int {
	numero := valor
	secuencia := 0
	return func() int {
		secuencia++
		return numero * secuencia
	}
}

// Cada vez que yo llame a MiTabla, este va a llamar a una función
func LlamarClosure() {
	tabladel := 2
	Mitabla := tabla(tabladel)
	for i := 1; i <= 10; i++ {
		fmt.Println(Mitabla())
	}
}
