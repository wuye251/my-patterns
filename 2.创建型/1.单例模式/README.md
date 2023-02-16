> 程序运行的期间，期望**某个类/配置文件/驱动**只需要初始化一次, 其他方法调用该类方法/属性都通过这一个实例，而不是重复创建/销毁,  为达到这个目的做的设计，称为单例模式



### 解决问题

1. 资源访问冲突（log文件写入覆盖,需要加互斥锁, 但如果不是同一个对象，每个调用都有各自的锁，不会互相互斥，所以单例就可以解决这种问题）
2. 配置文件(加载到内存后，应当只有一份)

### 如何实现
- 创建时线程安全
- 是否可以延迟加载
- 考虑getInstance()性能是否高

### 分类
> 单例模式分为两种实现：懒汉式、饿汉式   顾名思义懒汉式是**只有用到**这个配置文件/某个单例类/驱动时，**才创建**实例; 而饿汉式是在程序**刚启动时便已经创建**了实例，**不管是否会有方法来使用**
>
>  懒汉式是被动的，饿汉式是主动的。
#### 一、懒汉式
> 定义：只有使用该类时，才会实例化。
> 
> 优点：懒加载，节省未使用的单例的内存开销
> 
> 缺点：1. 首次使用速度慢  2.如果有问题(如oom 问题延后)
```go
type single struct{}

var singleInstance *single

func GetInstance() *single {
    // 只有在调用时才进行实例化
	if singleInstance == nil {
        singleInstance = &single{}
		fmt.Println("single instance already created.")
	}
	return singleInstance
}
```
#### 二、饿汉式
> 定义：在类加载期间，就已经将实例初始化好了。
> 
优点：
1. 如果有问题，将会放在程序启动时就暴露出来 
2. 首次/非首次访问速度一样 
3. 线程安全

缺点：
1. 首次启动时慢
2. 不支持延迟加载
```go
type single struct{}

var singleInstance *single

func init() {
    singleInstance = &single{}
}

func GetInstance() *single {
	return singleInstance
}
```
#### 三、双重检测(懒汉式)
> 多线程模型下，会出现多个线程同时初始化实例，如果不加锁限制，会导致多次实例化
```go
package singleton

import (
	"fmt"
	"sync"
)

var lock = &sync.Mutex{}

type single struct{}

var singleInstance *single

func GetInstance() *single {
	if singleInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		// 这里额外加一层判断是因为 首次初始化时，同时有多个协程进入，此时都会进入锁等待，等第一个协程释放锁之后，其他协程会在该流程继续执行，此时如果没有二次判断，会导致第一个协程创建的实例被覆盖
		if singleInstance == nil {
			fmt.Println("create new single.")
			singleInstance = &single{}
		}

	} else {
		fmt.Println("single instance already created.")
	}

	return singleInstance
}
```
