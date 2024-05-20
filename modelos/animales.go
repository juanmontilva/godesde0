package modelos

type Animales struct {
	Respira  bool
	Especies string
}

func (a *Animales) Respirar() {
	a.Respira = true
}

func (a *Animales) Especie() string {
	return "Felino"
}

func (a *Animales) EsCarnivoro() bool {
	return true
}

func (a *Animales) Comer() string {
	return "carne"
}

func (a *Animales) EstaVivo() bool {
	return true
}
