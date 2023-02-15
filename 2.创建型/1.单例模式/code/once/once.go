package once

import (
	"fmt"
	"sync"
)

var once sync.Once

type single struct {
}

var singleInstance *single

func GetInstance() *single {
	if singleInstance == nil {
		once.Do(func() {
			fmt.Println("create new single.")
			singleInstance = &single{}
		})
	} else {
		fmt.Println("single instance already created.")
	}

	return singleInstance
}
