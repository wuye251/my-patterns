package my_builder

import "fmt"

type PersonFatBuilder struct {
	head     int
	body     int
	armLeft  int
	armRight int
	legLeft  int
	legRight int
}

func (fat *PersonFatBuilder) BuildHead() {
	fat.head = 10
	fmt.Println("fat build head")
}

func (fat *PersonFatBuilder) BuildBody() {
	fat.body = 10
	fmt.Println("fat build body")
}

func (fat *PersonFatBuilder) BuildArmLeft() {
	fat.armLeft = 10
	fmt.Println("fat build left arm")
}
func (fat *PersonFatBuilder) BuildArmRight() {
	fat.armRight = 10
	fmt.Println("fat build right arm")
}
func (fat *PersonFatBuilder) BuildLegLeft() {
	fat.legLeft = 10
	fmt.Println("fat build left leg")
}
func (fat *PersonFatBuilder) BuildLegRight() {
	fat.legRight = 10
	fmt.Println("fat build right leg")
}
