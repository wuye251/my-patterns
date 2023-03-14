package my_observer

import "fmt"

type IObserver interface {
	Update(string)
}

type Observer1 struct{ Name string }

func (obs1 *Observer1) Update(message string) {
	fmt.Printf("obs1 %s update message is %s\n", obs1.Name, message)
}

type Observer2 struct {
	Attr string
}

func (obs2 *Observer2) Update(message string) {
	fmt.Printf("obs1 %s update  messageg is %s\n", obs2.Attr, message)
}
