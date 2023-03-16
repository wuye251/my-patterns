package decorator

type Cheese struct {
	Pizza IPizza
}

func (cheese *Cheese) GetPrice() int {
	return cheese.Pizza.GetPrice() + 5
}
