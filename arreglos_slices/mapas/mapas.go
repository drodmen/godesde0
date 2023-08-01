package mapas

import "fmt"

func MostrarMapas() {
	paises := make(map[string]string)
	// muestra el mapa antes de tener valores
	fmt.Println(paises)

	//introduciendo valores
	paises["México"] = "DF"
	paises["Argentina"] = "Buenos Aires"
	// muestran todos los elementos del mapa
	fmt.Println(paises)
	// muestra el elemento del mapa que está consignado/unido a la clave "Argentina"
	fmt.Println(paises["Argentina"])

	// cómo crear un mapa directamente sin el make
	campeonato := map[string]int{
		"Barcelona":    39,
		"Real Madrid":  50,
		"Chivas":       37,
		"Boca Juniors": 25,
	}

	fmt.Println(campeonato)

	// recorrer mapa
	for equipo, puntuaje := range campeonato {
		fmt.Printf("Equipo %s, tiene una puntuación de %d \n", equipo, puntuaje)
	}

	// quitar elementos de un mapa
	delete(campeonato, "Barcelona")
	fmt.Println(campeonato)

	//buscar por clave
	puntuaje, existe := campeonato["Juventus"]
	fmt.Printf("La puntuación obtenida es %d, y el equipo existe = %t \n", puntuaje, existe)

	puntuaje, existe2 := campeonato["Real Madrid"]
	fmt.Printf("La puntuación obtenida es %d, y el equipo existe = %t \n", puntuaje, existe2)

}
