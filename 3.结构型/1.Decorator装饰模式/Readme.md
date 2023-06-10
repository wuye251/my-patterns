> 在不修改原始实现类基础上，动态扩展额外的一些功能，符合开闭原则

## 示例

一般去买披萨，在一张披萨的基础上可以选择多加芝士，多加肉酱等多种选择，不同披萨所加可以的配料也是不同的，如果是蔬菜披萨，是不可以加肉酱的， 而且店家也时常会更新(上架/下架配料)， 如何做到程序设计考虑到这些，易于扩展，但又不影响现有功能， 本节的装饰模式便使用这个场景

``` go
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
	return decorator.Pizza.GetPrice()
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
	tomatoDecorator := NewTomatoDecorator(decorator)
	// 多加一份芝士
	cheeseDecorator := NewCheeseDecorator(tomatoDecorator)

	fmt.Printf("get pizza with tomato and cheese price is %d\n", cheeseDecorator.GetPrice())

}
```

### 实体和装饰器共用的interface

**装饰器的关键，需要实体和装饰器实现同一个接口,  意思是你可以把装饰器当成一个实体**， 这样做的好处是多个装饰器进行装饰时， 后面的装饰器可以使用上面的装饰器进行进一步的功能扩展，

不太清楚的可以看下上面代码中main()方法中 `tomatoDecorator`装饰后的披萨返回的实例，因为装饰器可以当成实体，所以下面的`CheeseDecorator`才可以在`tomatorDecorator`装饰后的基础上进一步进行扩展装饰

## UML

![image-20230611004244003](http://img.hahagblog.com/local/image-20230611004244003.png)

## 继承和装饰的区别

1. 装饰模式是在不改变原有对象的基础上，将功能附加到对象上

2. 继承是改变了原有对象的属性和方法

3. 装饰模式是动态的，继承是静态的

4. 装饰模式是多个装饰对象的组合，继承是单一的

5. 装饰模式是一种对象结构型模式

## 优点

1. 装饰器模式是继承关系的一个替代方案, 可以轻量级的扩展被装饰对象的功能

2. 装饰器模式通过使用不同的装饰类以及这些装饰类的排列组合, 可以实现不同的效果

3. 装饰器模式完全遵守开闭原则

## 缺点

1. 装饰器模式会增加许多子类, 一定程度上增加了系统的复杂性

2. 装饰器模式会引入许多的小对象, 大量小对象占据内存, 一定程度上影响程序的性能

3. 装饰器模式采用动态组合的方式, 并没有形成一套静态的组合关系, 复杂度较高

## 适用场景

1. 需要扩展一个类的功能或给一个类添加附加职责

2. 需要动态的给一个对象添加功能，这些功能可以再动态的撤销

3. 需要增加由一些基本功能的排列组合而产生的非常大量的功能，从而使继承关系变得不现实

4. 当不能采用生成子类的方法进行扩充时

## 参考

- [《B站刘丹冰--装饰模式》](https://www.bilibili.com/video/BV1Cg411e7Xi/?spm_id_from=333.788&vd_source=f53bb49fb78a32947a9360dd16a1cf58)
- 《Head First 第三篇装饰模式 P79》
- 《大话设计模式 -- 第六章装饰模式 P44》
