/* package main

import "fmt"

func add() func(y int) int {
	var x int
	return func(y int) int {
		x += y
		return x
	}
}

func cal(base int) (func(int) int, func(int) int) {
	add := func(i int) int {
		base += i
		return base
	}
	sub := func(i int) int {
		base -= i
		return base
	}
	return add, sub
}

func main() {
	add, sub := cal(100)
	fmt.Printf("add(100): %v\n", add(100))
	fmt.Printf("add(200): %v\n", add(200))
	fmt.Printf("sub(10): %v\n", sub(10))
	add, sub = cal(20)
	fmt.Printf("add(0): %v\n", add(0))
	// f := add()
	// fmt.Printf("f(10): %v\n", f(10))
	// fmt.Printf("f(20): %v\n", f(20))
	// fmt.Printf("f(30): %v\n", f(30))
	// f = add()
	// fmt.Printf("f(10): %v\n", f(10))
}
*/