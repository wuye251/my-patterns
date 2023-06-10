package main

import "fmt"

// 实体类抽象
type IPizza interface {
	GetPrice() int
}

// 实体类
type MeatPizza struct{}

// 肉酱披萨的最基本的实体价格
func (meat *MeatPizza) GetPrice() int {
	fmt.Println("get meat pizza")
	return 15
}

// 素披萨的最基本的实体
type VeggPizza struct{}

func (vegg *VeggPizza) GetPrice() int {
	fmt.Println("get vegg pizza")
	return 10
}

// 装饰器的基本对象 保存当前实体类/装饰器的实例
type Decorator struct {
	Pizza IPizza
}

func NewDecorator(pizza IPizza) IPizza {
	return &Decorator{Pizza: pizza}
}

// 基础类中不做任何操作, 只负责需要被装饰的对象保存起来  让当前装饰器对象使用
func (decorator *Decorator) GetPrice() int {
	return 0
}

// 添加西红柿配料的装饰器
type TomatoDecorator struct {
	Decorator
}

func NewTomatoDecorator(pizza IPizza) IPizza {
	return &TomatoDecorator{Decorator: Decorator{Pizza: pizza}}
}

func (tomato *TomatoDecorator) GetPrice() int {
	fmt.Println("add tomato to pizza")
	return tomato.Pizza.GetPrice() + 2
}

// 添加芝士配料的装饰器
type CheeseDecorator struct {
	Decorator
}

func NewCheeseDecorator(pizza IPizza) IPizza {
	return &CheeseDecorator{Decorator: Decorator{Pizza: pizza}}
}

func (cheese *CheeseDecorator) GetPrice() int {
	fmt.Println("add cheese to pizza")
	return cheese.Pizza.GetPrice() + 3
}

func main() {
	// 顾客想要一个肉酱披萨
	pizza := &MeatPizza{}

	// 开始装饰这个披萨
	decorator := NewDecorator(pizza)
	// 多加一份西红柿
	tomato := NewTomatoDecorator(decorator)
	// 多加一份芝士
	cheese := NewCheeseDecorator(tomato)

	fmt.Printf("get pizza with tomato and cheese price is %d\n", cheese.GetPrice())

}
