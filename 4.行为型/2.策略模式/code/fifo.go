package my_strategy

import "fmt"

type FIFOPop struct {
}

func (fifo FIFOPop) pop(cache *Cache) {
	fmt.Println("fifo pop")
}
