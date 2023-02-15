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
