package files

import (
	"bufio"
	"fmt" //manejo de archivos
	"os"

	"github.com/drodmen/godesde0/ejercicios"
)

var fileName string = "./files/txt/tabla.txt"

func GrabaTabla() {
	var texto string = ejercicios.TablasMultiplicar()
	archivo, err := os.Create(fileName) //Create te crea un archivo, si existe, lo reemplaza
	if err != nil {
		fmt.Println("Error al crear el archivo" + err.Error())
		return //muestra mensaje y se sale de la función
	}

	fmt.Fprintln(archivo, texto) //hace un archivo del texto
	archivo.Close()              //se cierra el archivo que se ha creado
}

// Concatena tablas de multiplicar
func SumaTabla() {
	var texto string = ejercicios.TablasMultiplicar()
	if !Append(fileName, texto) { //recibe un nombre de archivo (con su ruta) y el texto que va a introducir en el archivo
		fmt.Println("Error al concatenar el contenido")
	}
}

func Append(filen string, texto string) bool {
	arch, err := os.OpenFile(fileName, os.O_WRONLY|os.O_APPEND, 0644) // Con OpenFile, abre el archivo y lo abre para hacer concatenación al final del archivo
	// con os.O_WRONLY|os.O_APPEND enviamos dos variables, el 0644 lo que da es permiso de escritura al que es el dueño del archivo, el resto solo puede leer
	if err != nil {
		fmt.Println("Error durante el Append" + err.Error())
		return false
	}

	_, err = arch.WriteString(texto) //grabar lo que sea en el archivo
	if err != nil {
		fmt.Println("Error durante el Append" + err.Error())
		return false
	}

	arch.Close()
	return true
}

func LeoArchivo() {
	archivo, err := os.Open(fileName) //lo único que hace es lectura
	if err != nil {
		fmt.Println("Hubo un error leyendo el archivo" + err.Error())
		return
	}
	scanner := bufio.NewScanner(archivo)
	for scanner.Scan() {
		registro := scanner.Text() //devuelve cada línea del archivo
		fmt.Println("> " + registro)
	}
	archivo.Close()
}
