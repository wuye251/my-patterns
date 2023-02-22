package my_strategy

type Cache struct {
	storage     map[string]string
	popAlgo     PopAlgo
	capacity    int
	maxCapacity int
}

func InintCache(popAlgo PopAlgo) *Cache {
	return &Cache{
		popAlgo:     popAlgo,
		maxCapacity: 3,
		capacity:    0,
		storage:     make(map[string]string),
	}
}

func (cache *Cache) SetPopAlgo(popAlgo PopAlgo) {
	cache.popAlgo = popAlgo
}

func (cache *Cache) Add(key, value string) {
	if cache.capacity == cache.maxCapacity {
		cache.Pop()
	}
	cache.capacity++
	cache.storage[key] = value
}

func (cache *Cache) Get(key string) {
	delete(cache.storage, key)
}

func (cache *Cache) Pop() {
	cache.popAlgo.pop(cache)
	cache.capacity--
}
