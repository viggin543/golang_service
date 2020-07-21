package tests

import (
	"sync"
	"testing"
)

type compute struct {
	sync.RWMutex // this is extension class !!
	response map[int]int
}


func fib(i int) int {
	if i == 0 {
		return 0
	}
	if i == 1 {
		return 1
	}
	return fib(i-1) + fib(i-2)
}

func (c *compute) workerFib(bufferedChan chan int, wg *sync.WaitGroup) {
	for i := range bufferedChan {
		result := fib(i)
		c.Lock()
		c.response[i] = result
		c.Unlock()
	}
	wg.Done()
}

func (c *compute) parallelFibonacci(count int) {
	bufferedChan := make(chan int, 100)
	wg := &sync.WaitGroup{}
	for i := 0; i < 2; i++ {
		wg.Add(1)
		go c.workerFib(bufferedChan, wg)
	}
	for i := 0; i < count; i++ {
		bufferedChan <- i
	}
	close(bufferedChan)
	wg.Wait()
}

func TestFib(t *testing.T) {
	c := compute{response: make(map[int]int, 40)}
	c.parallelFibonacci(40)
}
