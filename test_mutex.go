/* package main

import (
	"fmt"
	"sync"
)

var c = make(chan int, 5)

var wg sync.WaitGroup

func add() {
	defer wg.Done()
	c <- 1
	// fmt.Printf("add:1\n")
}
func sub() {
	defer wg.Done()
	<-c
	//fmt.Printf("sub:1\n")
}

func main() {
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go add()
		wg.Add(1)
		go sub()
	}
	//time.Sleep(time.Second)
	wg.Wait()
	fmt.Printf("c: %v\n", len(c))
}
*/
/* //原子操作并不能保证同步
package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var counter int32 = 100

var wg sync.WaitGroup

func main() {
	for i := 0; i < 100; i++ {
		go func() {
			wg.Add(1)
			atomic.AddInt32(&counter, 1)
			wg.Done()
		}()
		go func() {
			wg.Add(1)
			atomic.AddInt32(&counter, -1)
			wg.Done()
		}()
	}
	//time.Sleep(time.Second)
	wg.Wait()
	fmt.Printf("counter: %v\n", counter)
}
*/
/* package main

import (
	"fmt"
	"sync"
	"time"
)

var i int = 100

var wg sync.WaitGroup

var lock sync.Mutex

func add() {
	// lock.Lock()
	defer wg.Done()
	i += 1
	fmt.Printf("i++: %v\n", i)
	time.Sleep(time.Millisecond * 20)
	// lock.Unlock()
}

func sub() {
	// lock.Lock()
	wg.Done()
	i -= 1
	fmt.Printf("i--: %v\n", i)
	time.Sleep(time.Millisecond * 500)
	// lock.Unlock()
}

func main() {
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go add()
		wg.Add(1)
		go sub()
	}
	wg.Wait()
	fmt.Printf("end i: %v\n", i)
}
*/
