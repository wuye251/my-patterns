package my_strategy_test

import (
	my_strategy "my-strategy"
	"testing"
)

// command `go test -timeout 30s -run ^TestCache$ my-strategy`
func TestCache(t *testing.T) {
	fifo := my_strategy.FIFOPop{}
	cache := my_strategy.InintCache(&fifo)

	cache.Add("a", "1")
	cache.Add("b", "3")
	cache.Add("c", "1")
	cache.Add("d", "4")

	lfu := my_strategy.LFUPop{}
	cache.SetPopAlgo(lfu)
	cache.Add("e", "5")

	lru := my_strategy.LRUPop{}
	cache.SetPopAlgo(lru)
	cache.Add("f", "6")
}
