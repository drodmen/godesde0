package interfaces

type Humano interface {
	//solo se coloca definición de funciones
	Respirar()
	Pensar()
	Comer()
	Sexo() string
	EstaVivo() bool
}
