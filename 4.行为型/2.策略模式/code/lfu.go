package my_strategy

import "fmt"

type LFUPop struct {
}

func (lfu LFUPop) pop(cache *Cache) {
	fmt.Println("lfu pop")
}
