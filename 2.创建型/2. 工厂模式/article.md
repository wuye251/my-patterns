# 简单工厂
> 将 `业务逻辑->基础方法`的依赖关系, 拆分为 `业务逻辑->实例创建->基础方法`,  后面添加实例(基础方法)类型时,不需要修改业务逻辑，符合单一职责

## 为什么需要简单工厂

举个例子来说, 如果我们要接入阿里oss,并且预计后续会有私有部署时使用的[minio](https://github.com/minio/minio)(本地搭建的oss服务)或者其他厂商的oss服务

这里我们先写第一版接入阿里oss

```go
package main

import (
	"fmt"
	"io"
)

type AliOSS struct {
	SecretId  string
	SecretKey string
	Bucket    string
}

// UploadFile 上传文件到阿里oss
func (AliOSS) UploadFile(localFilePath, objectKey string) error {
	fmt.Println("alioss upload file")
	return nil
}

// GetFile 下载/获取阿里oss文件
func (AliOSS) GetFile(objectKey string) (io.Reader, error) {
	fmt.Println("alioss get file")
	return nil, nil
}

// GetSignUrl 获取阿里oss临时访问地址
func (AliOSS) GetSignUrl(objectKey string) (string, error) {
	fmt.Println("alioss get sign url")
	return "", nil
}

func main() {
	// 创建逻辑
	alioss := AliOSS{
		SecretId:  "alioss_secret_id",
		SecretKey: "alioss_secret_key",
		Bucket:    "alioss_env_buckect_name",
	}

	// 业务逻辑
	alioss.UploadFile("/tmp/test.txt", "tmp/test.txt")
	alioss.GetFile("tmp/test.txt")
	alioss.GetSignUrl("tmp/test.txt")
}

```

我们第一版上线啦, 这样看上去还是很nice的

····· 时间流逝, 一个月过去了, 半年过去了

产品:"有个私有定制的项目,需要部署到客户内部机器中, 数据不能存到外面, 所以相关地方需要支持本地存储"

这时发现了oss使用的阿里oss,这需要修改成本地的, 于是开始调研哪些成熟的oss库支持本地搭建存储, 偶然间你突然发现了minio的存在, 可以本地搭建，速度也不赖, 该有的功能也全都支持, 这不正合需求嘛, 在你熟悉了提供的官方文档之后,果断撸起袖子加油干, 争取早日上线给客户使用

······ 时光荏苒, 一个月过去了, 私有部署的所有相关修改完成了, 关于你的oss代码成了这样:

```go
package main

import (
	"fmt"
	"io"
	"math/rand"
)

type AliOSS struct {
	EndPoint  string // 服务地址
	SecretId  string // 秘钥id
	SecretKey string // 秘钥key
	Bucket    string // 桶
}

// UploadFile 上传文件到阿里oss
func (AliOSS) UploadFile(localFilePath, objectKey string) error {
	fmt.Println("alioss upload file")
	return nil
}

// GetFile 下载/获取阿里oss文件
func (AliOSS) GetFile(objectKey string) (io.Reader, error) {
	fmt.Println("alioss get file")
	return nil, nil
}

// GetSignUrl 获取阿里oss临时访问地址
func (AliOSS) GetSignUrl(objectKey string) (string, error) {
	fmt.Println("alioss get sign url")
	return "", nil
}

type Minio struct {
	EndPoint string // 服务地址
	User     string // 用户
	Password string // 密码
	Bucket   string // 桶
}

// UploadFile 上传文件到minio
func (Minio) UploadFile(localFilePath, objectKey string) error {
	fmt.Println("minio upload file")
	return nil
}

// GetFile 下载/获取minio文件
func (Minio) GetFile(objectKey string) (io.Reader, error) {
	fmt.Println("minio get file")
	return nil, nil
}

// GetSignUrl 获取minio临时访问地址
func (Minio) GetSignUrl(objectKey string) (string, error) {
	fmt.Println("minio get sign url")
	return "", nil
}

func main() {
	// 获取配置文件判断是否私有部署
	isPrivate := rand.Intn(2)
	if isPrivate == 1 {
		minio := Minio{
			EndPoint: "minio_endpoint",
			User:     "minio_user",
			Password: "minio_password",
			Bucket:   "minio_env_bucket",
		}
		// 业务逻辑
		minio.UploadFile("/tmp/test.txt", "tmp/test.txt")
		minio.GetFile("tmp/test.txt")
		minio.GetSignUrl("tmp/test.txt")
	} else {
		alioss := AliOSS{
			EndPoint:  "alioss_endpoint",
			SecretId:  "alioss_secret_id",
			SecretKey: "alioss_secret_key",
			Bucket:    "alioss_env_buckect_name",
		}
		// 业务逻辑
		alioss.UploadFile("/tmp/test.txt", "tmp/test.txt")
		alioss.GetFile("tmp/test.txt")
		alioss.GetSignUrl("tmp/test.txt")
	}
}

```

这时就没有上次那么nice了,  main方法大了很多, 包含了根据是否私有化走不同oss逻辑, 既包含了创建的内容，也包含了业务逻辑的内容, 这还只是第二版,如果后期再迭代,那么main方法中的代码和逻辑会更庞杂，该如何优化呢?

- 将两个分支逻辑拆开封装成一个方法
  但是会有两个问题
  - 创建和业务逻辑还是耦合的
  - 具有重复的业务逻辑代码

那如何解决上面两个问题呢? 将创建和业务逻辑拆分,解耦合,   还有减少重复的业务逻辑代码

- 增加一个代理方法，里面根据不同if/else条件, 返回不同的实例类，然后根据接口提供的方法进行业务逻辑的实现, **降低了创建和业务逻辑的耦合性**
- 声明一个oss接口，指定有哪些功能方法，main方法中使用实现接口的实例类方法, 这样可以**减少业务逻辑代码**

  秘籍来了, 请用心感受下面代码对于上面两个问题的处理思想

```go
package main

import (
	"fmt"
	"io"
	"math/rand"
)

// oss 接口指定实现类需要提供哪些能力
type OSS interface {
	UploadFile(localFilePath, objectKey string) error
	GetFile(objectKey string) (io.Reader, error)
	GetSignUrl(objectKey string) (string, error)
}

type AliOSS struct {
	EndPoint  string // 服务地址
	SecretId  string // 秘钥id
	SecretKey string // 秘钥key
	Bucket    string // 桶
}

// UploadFile 上传文件到阿里oss
func (AliOSS) UploadFile(localFilePath, objectKey string) error {
	fmt.Println("alioss upload file")
	return nil
}

// GetFile 下载/获取阿里oss文件
func (AliOSS) GetFile(objectKey string) (io.Reader, error) {
	fmt.Println("alioss get file")
	return nil, nil
}

// GetSignUrl 获取阿里oss临时访问地址
func (AliOSS) GetSignUrl(objectKey string) (string, error) {
	fmt.Println("alioss get sign url")
	return "", nil
}

type Minio struct {
	EndPoint string // 服务地址
	User     string // 用户
	Password string // 密码
	Bucket   string // 桶
}

// UploadFile 上传文件到minio
func (Minio) UploadFile(localFilePath, objectKey string) error {
	fmt.Println("minio upload file")
	return nil
}

// GetFile 下载/获取minio文件
func (Minio) GetFile(objectKey string) (io.Reader, error) {
	fmt.Println("minio get file")
	return nil, nil
}

// GetSignUrl 获取minio临时访问地址
func (Minio) GetSignUrl(objectKey string) (string, error) {
	fmt.Println("minio get sign url")
	return "", nil
}

// 生成实例的代理工厂
type Factory struct{}

// 生成实例的不同判断逻辑
func (f Factory) CreateOSS() OSS {
	// 获取配置文件判断是否私有部署
	isPrivate := rand.Intn(2)
	switch isPrivate {
	case 1:
		return f.createMinio()
	default:
		return f.createAliOSS()
	}
}

// createMinio 生成minio的实例
func (Factory) createMinio() *Minio {
	// (省略)获取配置文件中minio的相关配置
	return &Minio{
		EndPoint: "minio_endpoint",
		User:     "minio_user",
		Password: "minio_password",
		Bucket:   "minio_env_bucket",
	}
}

// createAliOSS 生成alioss的实例
func (Factory) createAliOSS() *AliOSS {
	return &AliOSS{
		EndPoint:  "alioss_endpoint",
		SecretId:  "alioss_secret_id",
		SecretKey: "alioss_secret_key",
		Bucket:    "alioss_env_buckect_name",
	}
}

func main() {
	// 创建实例
	oss := Factory{}.CreateOSS()
	// 业务逻辑
	oss.UploadFile("/tmp/test.txt", "tmp/test.txt")
	oss.GetFile("tmp/test.txt")
	oss.GetSignUrl("tmp/test.txt")
}


```

如果后期如果再加其他oss, 我们只需要实现oss的接口方法, 并且添加到创建代理工厂中，而不影响业务逻辑中的代码了,是不是很nice呢!

## 简单工厂的思想

这便是简单工厂模式的思想

- 提供一个工厂(代理)，完成创建工作,   减少创建和业务逻辑代码的耦合性
- 通过接口定义, 业务逻辑面相接口实现功能逻辑, 增加其他类型不影响业务代码， 只需要实现接口的方法以及注册到工厂中即可

因为在上面已经详细推导过过程，不再叙述工厂模式了, 只要记得  `接口+工厂(代理工厂)解耦创建和业务逻辑`就行, 提供一个类图:

## UML
![image-20230525223622888.png](http://img.hahagblog.com/FgIdkYHMm_eGBiTRMeretb_K8Jed)

## 什么时候用

例子围绕oss一个具体场景来使用简单工厂,  在日常工作中遇到的肯定是五花八门的,  提供一个技巧，什么时候需要考虑用简单工厂
简单工厂属于**创建型**, 围绕创建解决问题, 故而重点应该放在创建上, 如果有相同功能的不同类型(minio、alioss都是oss上传下载)场景下, 可以将创建这个功能(逻辑) 改成一个简单工厂， 后面的业务逻辑不依赖不同实例产生不同分支

## 优缺点
**优点**

- 符合单一职责, 业务逻辑和创建功能拆分
- 面向接口,  业务逻辑针对oss的接口进行逻辑处理, 这样再添加类型的时候不影响业务逻辑

**缺点**

- 违反开闭原则, 增加一个类型，需要修改工厂内容
- 随着类型增加，工厂越来越庞大，类型实例也会增加,  增加代码复杂度

> 回顾简单工厂，我们将创建oss(产品)实例的代码解耦到工厂中进行，实现了业务逻辑和创建的解耦，也避免了代码的重复。但是有没有发现创建的工厂中还是存在了一堆的if/else或是switch判断语句，这些和创建实际是无关的，但是被耦合到工厂里面，若是后面再添加一种oss类型，实现产品的基础上，我们还需要修改工厂，因为他耦合了switch的分支判断， 很明显违反了开闭原则(对扩展是开放的，对修改是关闭的)。

# 工厂模式
## 工厂模式解决了什么

而工厂模式正是解决了简单模式中违反开闭原则的缺点，换句话来说 `工厂模式 =简单工厂+解决开闭原则`，着重将工厂中判断分支逻辑解耦， 工厂中只剩下创建实例的部分， 那么到底是如何解决的呢，看下根据简单工厂的例子进行的优化:

```go
package main

import (
	"fmt"
	"io"
	"math/rand"
)

// oss 接口指定实现类需要提供哪些能力
type OSS interface {
	UploadFile(localFilePath, objectKey string) error
	GetFile(objectKey string) (io.Reader, error)
	GetSignUrl(objectKey string) (string, error)
}

type AliOSS struct {
	EndPoint  string // 服务地址
	SecretId  string // 秘钥id
	SecretKey string // 秘钥key
	Bucket    string // 桶
}

// UploadFile 上传文件到阿里oss
func (AliOSS) UploadFile(localFilePath, objectKey string) error {
	fmt.Println("alioss upload file")
	return nil
}

// GetFile 下载/获取阿里oss文件
func (AliOSS) GetFile(objectKey string) (io.Reader, error) {
	fmt.Println("alioss get file")
	return nil, nil
}

// GetSignUrl 获取阿里oss临时访问地址
func (AliOSS) GetSignUrl(objectKey string) (string, error) {
	fmt.Println("alioss get sign url")
	return "", nil
}

type Minio struct {
	EndPoint string // 服务地址
	User     string // 用户
	Password string // 密码
	Bucket   string // 桶
}

// UploadFile 上传文件到minio
func (Minio) UploadFile(localFilePath, objectKey string) error {
	fmt.Println("minio upload file")
	return nil
}

// GetFile 下载/获取minio文件
func (Minio) GetFile(objectKey string) (io.Reader, error) {
	fmt.Println("minio get file")
	return nil, nil
}

// GetSignUrl 获取minio临时访问地址
func (Minio) GetSignUrl(objectKey string) (string, error) {
	fmt.Println("minio get sign url")
	return "", nil
}

// 生成实例抽象工厂
type Factory interface {
	CreateOSS() OSS
}

// 每个产品提供一个创建工厂
type AliOSSFactory struct{}

// CreateOSS  生成alioss 实例的工厂
func (AliOSSFactory) CreateOSS() *AliOSS {
	// (省略)获取配置文件中minio的相关配置
	return &AliOSS{
		EndPoint:  "alioss_endpoint",
		SecretId:  "alioss_secret_id",
		SecretKey: "alioss_secret_key",
		Bucket:    "alioss_env_buckect_name",
	}
}

type MinioFactory struct{}

// CreateOSS 生成minio的实例工厂
func (MinioFactory) CreateOSS() *Minio {
	// (省略)获取配置文件中minio的相关配置
	return &Minio{
		EndPoint: "minio_endpoint",
		User:     "minio_user",
		Password: "minio_password",
		Bucket:   "minio_env_bucket",
	}
}

func main() {
	// 获取配置文件判断是否私有部署
	var oss OSS
	isPrivate := rand.Intn(2)
	switch isPrivate {
	case 1:
		oss = MinioFactory{}.CreateOSS()
	default:
		oss = AliOSSFactory{}.CreateOSS()
	}

	// 业务逻辑
	oss.UploadFile("/tmp/test.txt", "tmp/test.txt")
	oss.GetFile("tmp/test.txt")
	oss.GetSignUrl("tmp/test.txt")
}
```

如上面代码所示，声明了一个抽象工厂，每个产品的创建工厂必须实现抽象工厂的创建方法， 但他们又不耦合于创建前的分支判断，单纯的是创建实例即可， 而分支判断逻辑则被修改到了main方法(客户端), 根据客户端的需求想要哪个产品，就可以跟对应的产品工厂要。

后续如果要添加一个新的oss，我们需要实现产品实力的同时，创建一个相应的创建工厂， 好处是不需要修改/影响现有的其他产品/创建工厂，  而改动的影响被迁移到了客户端中，完美解决了 简单工厂中违反开闭原则的缺点

>  此处我觉得还符合了依赖倒转的原则，改动被提到了客户端(经常修改的的地方)，而底层的方法(产品类、工厂创建类)不受影响，客户端依赖底层抽象，而不是具体的实现

## UML
![image-20230525235012856.png](http://img.hahagblog.com/FquHueChDONn6I3FvKWvRhmwhRva)

## 优缺点

**优点**

1. 符合开闭原则
2. 创建和使用解耦
3. 扩展性更好
4. 继承简单工厂的优点
   1. 单一职责

**缺点**

> 可以发现不断地优化解决，附加的设计也越来越多， 之前添加一个实例类型，只需要实现产品接口，再到工厂中加一个分支即可， 现在也需要添加一个工厂类来实现创建， 更复杂了些

# 抽象工厂
## 介绍

抽象工厂模式是工厂模式的升级版，它用来创建一组相关或者相互依赖的对象。

他与工厂模式的区别就在于，工厂模式针对的是一个产品等级结构；而抽象工厂模式则针对的是面向多个产品等级结构的。

在编程中，通常一个产品结构，表现为一个接口或者抽象类，也就是说，工厂模式提供的所有产品都是衍生自同一个接口或抽象类，

而抽象工厂模式所提供的产品则是衍生自不同的接口或抽象类。

例子如下:

```go
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

```
## UML
![image-20230527010540613.png](http://img.hahagblog.com/FgqzULjYa5O7QH13X09loctTkEW1)
## 如何更简单的理解抽象工厂

抽象工厂模式的实质是：提供接口，创建一系列相关或相互依赖的对象。

举个例子：我们去麦当劳点餐，点的是套餐，套餐里有汉堡、薯条、可乐，这些东西是相互依赖的，汉堡里有肉，薯条里有土豆，

可乐里有糖，这些东西是相互依赖的，如果没有其中的一个，那么这个套餐就不完整了。

但是，我们去麦当劳点餐的时候，不会去点汉堡、薯条、可乐，而是直接点套餐，因为套餐里包含了汉堡、薯条、可乐，这样更方便。

这里的套餐就是抽象工厂，汉堡、薯条、可乐就是抽象工厂里面的产品，而麦当劳就是抽象工厂的实现类。


## 优缺点

**优点**

1. 抽象工厂模式隔离了具体类的生成，使得客户并不需要知道什么被创建。由于这种隔离，更换一个具体工厂就变得相对容易，所有的具体工厂都实现了抽象工厂中定义的那些公共接口，因此只需改变具体工厂的实例，就可以在某种程度上改变整个软件系统的行为。

2. 当一个产品族中的多个对象被设计成一起工作时，它能够保证客户端始终只使用同一个产品族中的对象。

3. 增加新的产品族很方便，无须修改已有系统，符合“开闭原则”。

**缺点**

1. 在添加新的产品对象时，难以扩展抽象工厂来生产新种类的产品，这是因为在抽象工厂角色中规定了所有可能被创建的产品集合，要支持新种类的产品就意味着要对该接口进行扩展，而这将涉及到对抽象工厂角色及其所有子类的修改，显然会带来较大的不便。

2. 开闭原则的倾斜性（增加新的工厂和产品族容易，增加新的产品等级结构麻烦）。

## 使用场景

1. 一个系统不应当依赖于产品类实例如何被创建、组合和表达的细节，这对于所有形态的工厂模式都是重要的。

2. 系统中有多于一个的产品族，而每次只使用其中某一产品族。

3. 属于同一个产品族的产品将在一起使用，这一约束必须在系统的设计中体现出来。

4. 系统提供一个产品类的库，所有产品以同样的接口出现，从而使客户端不依赖于具体实

# 总结
## 产品和工厂的对应关系

- 简单工厂是 **一对多**的关系(一个工厂 对应多个产品)
- 工厂模式是 **一对一**的关系(每个产品分别对应各自一个工厂)
- 抽象工厂是 **多对一**的关系(多个产品对应一个工厂)

## 解决问题

- 简单工厂解决**创建和业务逻辑耦合，面向过程改为面向接口编程**

- 工厂模式**解决简单工厂违反开闭原则**，添加产品不修改现有工厂
- 抽象工厂添加产品等级结构, **一个工厂可组合多个产品，工厂模式只有一种产品**

## 一句话

- 简单工厂：创建和业务逻辑解耦，业务面相接口编程
- 工厂模式：一个产品一个工厂，符合开闭原则
- 抽象工厂：多个产品一个工厂


# 参考
- [b站刘丹冰 -- 工厂模式](https://www.bilibili.com/video/BV1we4y1H7NU/?spm_id_from=pageDriver&vd_source=f53bb49fb78a32947a9360dd16a1cf58)
- 《大话设计模式》--- 第1、8、15章
- 《Head First设计模式》 -- 第四篇工厂模式