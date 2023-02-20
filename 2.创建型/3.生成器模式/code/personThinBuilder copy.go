package my_builder

import "fmt"

type PersonThinBuilder struct {
	head     int
	body     int
	armLeft  int
	armRight int
	legLeft  int
	legRight int
}

func (thin *PersonThinBuilder) BuildHead() {
	thin.head = 1
	fmt.Println("thin build head")
}

func (thin *PersonThinBuilder) BuildBody() {
	thin.body = 1
	fmt.Println("thin build body")
}

func (thin *PersonThinBuilder) BuildArmLeft() {
	thin.armLeft = 1
	fmt.Println("thin build left arm")
}
func (thin *PersonThinBuilder) BuildArmRight() {
	thin.armRight = 1
	fmt.Println("thin build right arm")
}
func (thin *PersonThinBuilder) BuildLegLeft() {
	thin.legLeft = 1
	fmt.Println("thin build left leg")
}
func (thin *PersonThinBuilder) BuildLegRight() {
	thin.legRight = 1
	fmt.Println("thin build right leg")
}
