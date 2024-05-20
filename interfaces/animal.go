package interfaces

type Animal interface {
	Respirar()
	Comer() string
	EsCarnivoro() bool
	EstaVivo() bool
	Especie() string
}
