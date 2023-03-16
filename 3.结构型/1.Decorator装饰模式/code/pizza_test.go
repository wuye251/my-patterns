package decorator_test

import (
	"decorator"
	"fmt"
	"testing"
)

func TestPizza(t *testing.T) {
	pizza := decorator.Vegg{}
	fmt.Printf("get price is %d\n", pizza.GetPrice())

	pizzaWithTomato := &decorator.Tomato{
		Pizza: &pizza,
	}
	fmt.Printf("get pizza with tomato price is %d\n", pizzaWithTomato.GetPrice())

	pizzaWithTomatoAndCheese := &decorator.Tomato{
		Pizza: pizzaWithTomato,
	}
	fmt.Printf("get pizza with tomato and cheese price is %d\n", pizzaWithTomatoAndCheese.GetPrice())

}
