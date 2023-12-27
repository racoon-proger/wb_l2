package main

import "fmt"

type Cache struct {
	storage        map[string]string
	evectionalAlgo EvictionalAlgo
	capacity       int
	maxCapacity    int
}

func initCache(e EvictionalAlgo) *Cache {
	storage := make(map[string]string)
	return &Cache{
		storage:        storage,
		evectionalAlgo: e,
		capacity:       0,
		maxCapacity:    2,
	}
}

func (c *Cache) setEvictionalAlgo(e EvictionalAlgo) {
	c.evectionalAlgo = e

}

func (c *Cache) add(key, value string) {
	if c.capacity == c.maxCapacity {
		c.evict()
	}

	c.capacity++
	c.storage[key] = value
}

func (c *Cache) get(key string) {
	delete(c.storage, key)
}

func (c *Cache) evict() {
	c.evectionalAlgo.evict(c)
	c.capacity--
}

type EvictionalAlgo interface {
	evict(c *Cache)
}

type Fifo struct {
}

func (l *Fifo) evict(c *Cache) {
	fmt.Println("Evicting by fifo stratrgy")
}

type Lry struct {
}

func (l *Lry) evict(c *Cache) {
	fmt.Println("Evicting by lru stratrgy")
}

type Lfu struct {
}

func (l *Lfu) evict(c *Cache) {
	fmt.Println("Evicting by lfu stratrgy")
}

func main() {
	lfu := &Lfu{}
	cache := initCache(lfu)

	cache.add("a", "1")
	cache.add("b", "2")

	cache.add("c", "3")

	lru := &Lry{}
	cache.setEvictionalAlgo(lru)

	cache.add("d", "4")

	fifo := &Fifo{}
	cache.setEvictionalAlgo(fifo)
	cache.add("e", "5")
}
