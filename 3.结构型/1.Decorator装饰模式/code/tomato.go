package decorator

type Tomato struct {
	Pizza IPizza
}

func (tomato *Tomato) GetPrice() int {
	return tomato.Pizza.GetPrice() + 5
}
