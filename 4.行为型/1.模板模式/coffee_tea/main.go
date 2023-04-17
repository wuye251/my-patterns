package main

import "fmt"

// 抽象类声明有哪些方法步骤
type Beverage interface {
	BoilWater()
	Brew()
	PutInCup()
	AddThings()
}

// 模板基类串联抽象方法
type template struct {
	maker Beverage
}

func (t *template) Make() {
	if t == nil {
		return
	}
	t.maker.BoilWater()
	t.maker.Brew()
	t.maker.PutInCup()
	t.maker.AddThings()
}

// 具体实现
type Tea struct {
	template //继承模板的功能
}

func (tea *Tea) BoilWater() {
	fmt.Println("茶水煮到100度")
}
func (tea *Tea) Brew() {
	fmt.Println("冲泡茶水")
}
func (tea *Tea) PutInCup() {
	fmt.Println("茶水倒到杯中")
}
func (tea *Tea) AddThings() {
	fmt.Println("茶水加小料")
}

func MakeTea() *Tea {
	makerTea := Tea{}
	// .maker声明茶的接口  实现多态
	makerTea.maker = &makerTea
	return &makerTea
}

// 具体实现
type Coffee struct {
	template //继承模板的功能
}

func (Coffee *Coffee) BoilWater() {
	fmt.Println("咖啡煮到100度")
}
func (Coffee *Coffee) Brew() {
	fmt.Println("冲泡咖啡")
}
func (Coffee *Coffee) PutInCup() {
	fmt.Println("咖啡倒到杯中")
}
func (Coffee *Coffee) AddThings() {
	fmt.Println("咖啡加小料")
}

func MakeCoffee() *Coffee {
	makerCoffee := Coffee{}
	// .maker声明茶的接口  实现多态
	makerCoffee.maker = &makerCoffee
	return &makerCoffee
}
func main() {
	// 制作茶
	MakeTea().Make()
	fmt.Println("=====")
	// 制作咖啡
	MakeCoffee().Make()
}

/* output:
茶水煮到100度
冲泡茶水
茶水倒到杯中
茶水加小料
=====
咖啡煮到100度
冲泡咖啡
咖啡倒到杯中
咖啡加小料
*/
