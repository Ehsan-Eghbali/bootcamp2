package main

import (
	"fmt"
	"sync"
	"time"
)

type calculator struct {
	*sync.WaitGroup
	count int
}
type result struct {
	*sync.WaitGroup
	number int
}

func main() {
	rw := new(sync.WaitGroup)
	nc := &calculator{
		rw,
		0,
	}
	rw1 := new(sync.WaitGroup)
	r := &result{
		rw1,
		1,
	}
	for i := 0; i < 100000; i++ {
		nc.Add(1)

		go nc.CalAdd()

		for j := 0; j < 10; j++ {
			rw1.Add(1)
			go r.multiply(i)
			rw1.Wait()
		}
		nc.Wait()
	}
	nc.Wait()
	time.Sleep(5 * time.Second)
	fmt.Println(nc.count)

}

func (nc *calculator) CalAdd() {
	defer nc.Done()
	fmt.Println(nc.count)
	nc.count++
}

func (r *result) multiply(int2 int) {
	defer r.Done()
	fmt.Println(int2 * r.number)
}
