package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	// Ожидание завершения всех горутин
	var wg sync.WaitGroup
	for i := 1; i <= 5; i++ {
		wg.Add(1)

		go func(n int) {
			defer wg.Done()
			worker(n)
		}(i)
	}
	wg.Wait()

	fmt.Println("------------------------------")

	// Атомарные арифметические операции
	var ops uint64
	wg = sync.WaitGroup{}
	for i := 0; i < 50; i++ {
		wg.Add(1)

		go func() {
			for c := 0; c < 1000; c++ {
				atomic.AddUint64(&ops, 1)
			}
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("ops:", ops)

	fmt.Println("------------------------------")

	// Мьютексы
	c := Container{
		counters: map[string]int{"a": 0, "b": 0},
	}
	wg = sync.WaitGroup{}

	doIncrement := func(name string, n int) {
		for i := 0; i < n; i++ {
			c.inc(name)
		}
		wg.Done()
	}

	wg.Add(3)
	go doIncrement("a", 10000)
	go doIncrement("a", 10000)
	go doIncrement("b", 10000)

	wg.Wait()
	fmt.Println(c.counters)
}

func worker(id int) {
	fmt.Printf("Worker %d starting\n", id)
	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)
}

type Container struct {
	mu       sync.Mutex
	counters map[string]int
}

func (c *Container) inc(name string) {
	// Блокировка критической секции через мьютекс
	c.mu.Lock()
	defer c.mu.Unlock()
	c.counters[name]++
}
