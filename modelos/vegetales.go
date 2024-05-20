package modelos

type Vegetales struct {
}

func (Vege *Vegetales) ClasificacionVegetal() string {
	return "Verdura"
}

func (Vege *Vegetales) EstaVivo() bool {
	return true
}
