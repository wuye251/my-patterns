> 多个类之间具有相同功能的方法，部分方法逻辑完全相同，部分方法逻辑不同， 那么可以提供一个抽象方法声明他们所有的方法，  一个基类实现部分逻辑完全相同的方法，  部分逻辑不相同的方法分别在各自类中独自实现

## 例子

例如制作茶/咖啡， 他们都有相同步骤，只是内部逻辑/原料不同

1. 烧开水
2. 放入茶叶/咖啡 冲泡
3. 倒入杯中
4. 添加糖

对比代码，

1. 声明一个接口interface 声明有哪些步骤
2. 一个基类，串联这些**抽象**步骤
3. 定义咖啡/茶的实体类  实现每个步骤的具体内容

代码如下：

```go
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

```

## UML

![image-20230612230254253](http://img.hahagblog.com/local/image-20230612230254253.png)

## 总结

通过模板方法，我们能从`template`类编程中感受到面向接口编程的舒适，不论具体类是哪些，扩展/修改都不会影响template的内容，  可以想象一个场景：我们需要对接第三方的某个功能(如消息队列)

- template中我们写关于队列拉取/入队前后的业务逻辑， 而内部的具体拉取/入队的实现和具体用哪个服务商，我们不需要关系
- Beverage规定队列的具有的功能， 不管对接哪些第三方，我们只需在各自的文件中进行编写第三方关于队列的入队/出队操作即可

这样的解耦，不至于让第三方的参数构建、配置初始化等操作代码出现在业务逻辑中。

## 优点

1. 封装不变部分，扩展可变部分

2. 提取公共代码，便于维护

3. 行为由父类控制，子类实现

## 适用场景

1. 一次性实现一个算法的不变的部分，并将可变的行为留给子类来实现(_对接第三方可变， 业务逻辑不变_)。

2. 各子类中公共的行为被提取出来并集中到一个公共父类中，从而避免代码重复。

3. 控制子类扩展。

