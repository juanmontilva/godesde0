package interfaces

type Humano interface {
	SerVivo
	Respirar()
	Pensar()
	Comer()
	Sexo() string
}
