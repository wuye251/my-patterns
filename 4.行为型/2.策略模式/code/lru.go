package my_strategy

import "fmt"

type LRUPop struct {
}

func (lru LRUPop) pop(cache *Cache) {
	fmt.Println("lru pop")
}
