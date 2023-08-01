package arreglos_slices

import "fmt"

var tablaS []int = []int{2, 5, 4} //crear un slice
var arreglo [10]int = [10]int{4, 52, 52, 23, 212, 11, 21, 45}

func MuestroSlices() {
	fmt.Println(tablaS)

	porcion := arreglo[3:]   //Slice creado desde el elemento 3 hasta el final
	porcion2 := arreglo[:4]  //Slice creado desde el elemento 0, hasta el 4
	porcion3 := arreglo[3:7] //Slice creado desde el elemento 3 hasta el 4

	fmt.Println(porcion)
	fmt.Println(porcion2)
	fmt.Println(porcion3)
}

//la capacidad es el espacio en memoria que reserva go, cuando se crea un slice.
func Capacidad() {
	elementos := make([]int, 5, 20) //permite hacer slice vacío dándole capacidad que se desee. Aquí se le dice que tendrá 5 elementos y se le da una capacidad de 20

	fmt.Printf("Largo %d, Capacidad %d", len(elementos), cap(elementos))

	nums := make([]int, 0, 0)
	for i := 0; i < 100; i++ {
		nums = append(nums, i)
	}

	fmt.Printf("Largo %d, Capacidad %d", len(nums), cap(nums))

}
