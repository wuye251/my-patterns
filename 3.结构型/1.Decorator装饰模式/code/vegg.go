package decorator

type Vegg struct {
	Pizza IPizza
}

func (vegg *Vegg) GetPrice() int {
	return 15
}
