package abstractfactory

// 抽象工厂
type AbstractFactory interface {
	// 创建产品A
	CreateProductA() ProductA
	// 创建产品B
	CreateProductB() ProductB
}

// 产品A
type ProductA interface {
	// 产品A的方法
	MethodA()
}

// 产品B
type ProductB interface {
	// 产品B的方法
	MethodB()
}

// 具体产品A1
type ConcreteProductA1 struct {
}

func (p *ConcreteProductA1) MethodA() {
	// 具体实现
}

// 具体产品A2
type ConcreteProductA2 struct {
}

func (p *ConcreteProductA2) MethodA() {
	// 具体实现
}

// 具体产品B1
type ConcreteProductB1 struct {
}

func (p *ConcreteProductB1) MethodB() {
	// 具体实现
}

// 具体产品B2
type ConcreteProductB2 struct {
}

func (p *ConcreteProductB2) MethodB() {
	// 具体实现
}

// 具体工厂1
type ConcreteFactory1 struct {
}

func (f *ConcreteFactory1) CreateProductA() ProductA {
	return &ConcreteProductA1{}
}


func (f *ConcreteFactory1) CreateProductB() ProductB {
	return &ConcreteProductB1{}
}

// 具体工厂2
type ConcreteFactory2 struct {
}

func (f *ConcreteFactory2) CreateProductA() ProductA {
	return &ConcreteProductA2{}
}

func (f *ConcreteFactory2) CreateProductB() ProductB {
	return &ConcreteProductB2{}
}

// 客户端
type Client struct {
}

func (c *Client) Create(factory AbstractFactory) {
	// 创建产品A
	productA := factory.CreateProductA()
	// 创建产品B
	productB := factory.CreateProductB()
	// do something
	productA.MethodA()
	productB.MethodB()
}

// 客户端调用
func main() {
	client := &Client{}
	// 创建工厂1
	factory1 := &ConcreteFactory1{}
	client.Create(factory1)
	// 创建工厂2
	factory2 := &ConcreteFactory2{}
	client.Create(factory2)
}
