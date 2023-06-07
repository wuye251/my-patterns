> 通过在客户端和服务端之间添加一层代理,  代理层在不影响其他端的情况下进行额外的一些操作， 但前提必须限制在代理层和服务端属于同一个抽象定义

## 角色

- Subject抽象主题 -- 真实主题和代理主题的共同接口
- RealSubject真实主题 -- 主题的真实显现对象
- Proxy代理 -- 实现抽象主题，并且包含真实主题, 可以在客户端调用真实主题前后执行一些操作(附加操作)

## 代码实例

```go
package main

// 抽象主题
type Subject interface {
	Do() string
}

// 真实主题
type RealSubject struct{}

func (RealSubject) Do() string {
	return "real"
}

// 代理主题
type Proxy struct {
	real RealSubject
}

func (p Proxy) Do() string {
	var res string
	res += "pre:"
	res += p.real.Do()
	res += ":after"
	return res
}

func main() {
	var sub Subject
	sub = Proxy{RealSubject{}}
	println(sub.Do())
}

```

## UML

![image-20230607221247833](http://img.hahagblog.com/local/image-20230607221247833.png)

## 优点

1. 你可以在客户端毫无察觉的情况下控制服务对象。

2. 如果客户端对于服务对象的生命周期没有特殊要求， 你可以对生命周期进行管理。

3. 单一职责原则。 你可以将对代理的各种功能进行分解。

4. 开闭原则。 你可以在不对服务或客户端做出修改的情况下创建新代理。

5. 你可以在客户端与服务对象之间放置代理， 以限制对服务对象的访问。

## 缺点

1. 代理模式会造成系统设计中类的数目增加
2. 在客户端和目标对象增加一个代理对象，会造成请求处理速度变慢
3. 增加系统的复杂度
4. 代理模式实现很简单，但是要为每一个服务都去写代理类，工作量太大，不易管理

## 应用场景

1. 远程代理
2. 保护代理
3. 缓冲代理
