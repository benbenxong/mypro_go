/* package main

import "fmt"

func test1() {
	var a1 [3]int
	var a2 [3]string
	fmt.Printf("a1: %T\n", a1)
	fmt.Printf("a2: %T\n", a2)
	fmt.Printf("a1: %v\n", a1)
	fmt.Printf("a2: %v\n", a2)
}

func test2() {
	a1 := [3]int{1, 2, 3}
	fmt.Printf("a1: %v\n", a1)
	var a2 = [...]string{"aa a", "bbb"}
	fmt.Printf("a2: %v\n", a2)
	fmt.Printf("len(a2): %v\n", len(a2))
}

func test3() {
	var a1 [3]int
	a1[0] = 100
	a1[2] = 200
	fmt.Printf("a1: %v\n", a1)
	for i := 0; i < len(a1); i++ {
		fmt.Printf("a1[%v]: %v\n", i, a1[i])
	}
	fmt.Printf("\"---------------\": %v\n", "---------------")
	for i, v := range a1 {
		fmt.Printf("a1[%v]: %v\n", i, v)
	}
}
func main() {
	//test1()
	//test2()
	test3()
}
*/