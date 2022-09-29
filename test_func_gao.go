/* package main

import "fmt"

func sayHello(name string) {
	fmt.Printf("Hello %s", name)
}

func test1(name string, f func(string)) {
	f(name)
}

func add(a int, b int) int {
	return a + b
}

func sub(a int, b int) int {
	return a - b
}

func cal(oper string) func(int, int) int {
	switch oper {
	case "+":
		return add
	case "-":
		return sub
	default:
		return nil
	}
}

func main() {
	//test1("tom", sayHello)
	// ff := cal("+")
	// fmt.Printf("ff(1, 2): %v\n", ff(1, 2))
	// ff = cal("-")
	// fmt.Printf("ff(4, 1): %v\n", ff(4, 1))
	c := func(a int, b int) int {
		return a + b
	}(1, 2)
	fmt.Printf("c: %v\n", c)
}
*/