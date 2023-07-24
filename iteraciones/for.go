package iteraciones

import (
	"fmt"
)

// Funciones públicas
func Iterar() {
	i := 0
	// bucle, mientras que i sea menor que 10
	for i < 10 {
		fmt.Println(i) // muestra por pantalla el valor de i
		i++            //se incrementa en 1 el valor de i
	}

	// otra forma de escribir el go
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	// si quiero que el for vaya en vez de sumando uno, sumando 5. Además, podemos contar el número de veces que itera
	iteracion := 1 // variable inicializada para iterar
	for i := 0; i < 25; i += 5 {
		fmt.Println(i)
		fmt.Println("Iteracion", iteracion) // muestro por pantalla la iteración en la que nos encontramos
		iteracion++                         // aumenta el valor de la iteración en 1, tras acabar el bucle
	}

	// si quiero que vaya bajando y que se salga del bucle
	// el for también vale para listas, mapas, slides, etc. Utilizando el for runch, que ya se verá o incluso accediendo a las filas de las BBDD
	for i := 50; i > 0; i -= 5 {
		if i == 20 {
			continue //muestra todos los números menos el 20, porque lo que le dice es que continue el bucle
		}

		if i == 10 {
			break //rompe el bucle en cuanto llegue a ese valor
		}
		fmt.Println(i)
	}

}
