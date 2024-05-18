package modelos

type Mujer struct {
	Hombre
	//este metodo en poo es herencia
	//se puede agregar tambien
	//agregar nuevas propiedades
	//ejemplo
	//EsMadre bool
}

func (m *Mujer) Respirar() {
	m.Respirando = true
}

func (m *Mujer) Comer() {
	m.Comiendo = true
}

func (m *Mujer) Pensar() {
	m.Pensando = true
}

func (m *Mujer) Sexo() string {
	return "Mujer"
}

func (m *Mujer) EstaVivo() bool {
	return true
}
