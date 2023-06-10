> 在不修改原始实现类基础上，动态扩展额外的一些功能，符合开闭原则

## 示例

一般去买披萨，在一张披萨的基础上可以选择多加芝士，多加肉酱等多种选择，不同披萨所加可以的配料也是不同的，如果是蔬菜披萨，是不可以加肉酱的， 而且店家也时常会更新(上架/下架配料)， 如何做到程序设计考虑到这些，易于扩展，但又不影响现有功能， 本节的装饰模式便使用这个场景

``` go
package main

import "fmt"

// 实体类和装饰器公用抽象
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

// 添加西红柿配料的装饰器
type TomatoDecorator struct {
	Pizza IPizza
}

func (tomato *TomatoDecorator) GetPrice() int {
	fmt.Println("add tomato to pizza")
	return tomato.Pizza.GetPrice() + 2
}

// 添加芝士配料的装饰器
type CheeseDecorator struct {
	Pizza IPizza
}

func (cheese *CheeseDecorator) GetPrice() int {
	fmt.Println("add cheese to pizza")
	return cheese.Pizza.GetPrice() + 3
}

func main() {
	// 顾客想要一个肉酱披萨
	var pizza IPizza
	pizza = &MeatPizza{}
	// 多加一份西红柿
	pizza = &TomatoDecorator{Pizza: pizza}
	// 多加一份芝士
	pizza = &CheeseDecorator{Pizza: pizza}
	fmt.Printf("get pizza with tomato and cheese price is %d\n", pizza.GetPrice())

}
```

### 实体和装饰器共用的interface

**装饰器的关键，需要实体和装饰器实现同一个接口,  意思是你可以把装饰器当成一个实体**， 这样做的好处是多个装饰器进行装饰时， 后面的装饰器可以使用上面的装饰器进行进一步的功能扩展，

不太清楚的可以看下上面代码中main()方法中 `tomatoDecorator`装饰后的披萨返回的实例，因为装饰器可以当成实体，所以下面的`CheeseDecorator`才可以在`tomatorDecorator`装饰后的基础上进一步进行扩展装饰

## UML



## 继承和装饰的区别

1. 装饰模式是在不改变原有对象的基础上，将功能附加到对象上

2. 继承是改变了原有对象的属性和方法

3. 装饰模式是动态的，继承是静态的

4. 装饰模式是多个装饰对象的组合，继承是单一的

5. 装饰模式是一种对象结构型模式













- 继承(静态扩展)
- 装饰模式 -- 组合(运行时动态扩展)



- 使用时注意装饰器的顺序
- 太多个装饰器，生成一个对象时需要创建 1+n(装饰器实例)，   可以结合`工厂/生成器`组合使用
