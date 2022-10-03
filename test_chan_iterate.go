/* package main

import "fmt"

var c = make(chan int, 5)

func main() {
	go func() {
		for i := 0; i < 3; i++ {
			c <- i
		}
		close(c)
	}()

	// for {
	// 	v, ok := <-c
	// 	if ok {
	// 		//fmt.Printf("len(c): %v\n", len(c))
	// 		fmt.Printf("v: %v\n", v)
	// 	} else {
	// 		break
	// 	}
	// }
	for v := range c {
		fmt.Printf("len(c): %v\n", len(c))
		fmt.Printf("v: %v\n", v)

	}
	fmt.Printf("len(c): %v\n", len(c))
}
*/