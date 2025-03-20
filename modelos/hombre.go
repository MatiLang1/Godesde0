package modelos

type Hombre struct {
	edad       int
	altura     float32
	peso       float32
	respirando bool
	pensando   bool
	comiendo   bool
	vivo       bool
}

func (h *Hombre) Respirar()    { h.respirando = true }
func (h *Hombre) Comer()       { h.comiendo = true }
func (h *Hombre) Pensar()      { h.pensando = true }
func (h *Hombre) Sexo() string { return "Hombre" }

//para que la estructura "Hombre" interprete la interfaz "Humano", ambas deben tener las mismas funciones
//Hombre se convierte en Humano (implementa esa interfaz) por tener las mismas funciones
