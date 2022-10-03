/* package main

import (
	"fmt"
	"runtime"
	"time"
)

func a() {
	for i := 0; i < 10; i++ {
		fmt.Printf("a i: %v\n", i)
		time.Sleep(time.Millisecond * 100)
	}
}
func b() {
	for i := 0; i < 10; i++ {
		fmt.Printf("b i: %v\n", i)
		time.Sleep(time.Millisecond * 100)
	}
}

func main() {
	fmt.Printf("runtime.NumCPU(): %v\n", runtime.NumCPU())
	// runtime.GOMAXPROCS(14)
	go a()
	go b()

	time.Sleep(time.Second)
}

/* package main

import (
	"fmt"
	"runtime"
	"time"
)

func show() {
	for i := 0; i < 10; i++ {
		fmt.Printf("i: %v\n", i)
		if i == 5 {
			runtime.Goexit()
		}
		// time.Sleep(time.Microsecond * 100)
	}
}
func main() {
	go show()
	// for i := 0; i < 2; i++ {
	// 	runtime.Gosched()
	// 	fmt.Printf("\"golang\": %v\n", "golang")
	// }
	time.Sleep(time.Second)
	fmt.Println("end...")
}
*/
 */