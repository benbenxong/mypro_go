/* package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func showMsg(i int) {
	defer wg.Done()
	fmt.Printf("i: %v\n", i)
	time.Sleep(time.Second)
}

func main() {
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go showMsg(i)

	}

	wg.Wait()

	fmt.Printf("end...")
}
*/