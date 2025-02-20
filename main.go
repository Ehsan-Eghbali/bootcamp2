package main

import (
	"fmt"
	"sync"
)

var count int

func main() {
	//mu := new(sync.Mutex)
	//for i := 0; i < 1000; i++ {
	//	go addNumber(mu)
	//}
	//fmt.Println(count)

	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(2)
		go addNumber(&wg)
		wg.Wait()
	}
	fmt.Println(count)
}

//	func addNumber(mu *sync.Mutex) {
//		mu.Lock()
//		//defer mu.Unlock()
//		count++
//		mu.Unlock()
//	}
func addNumber(wg *sync.WaitGroup) {
	defer wg.Done()
	count++
}
