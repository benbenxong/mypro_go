/* package main

import "fmt"

func test1() {
	var s1 []int
	var s2 []string
	fmt.Printf("s1: %v\n", s1)
	fmt.Printf("s2: %v\n", s2)
}

func test2() {
	var s2 = make([]int, 2)
	fmt.Printf("s2: %v\n", s2)

}

func test3() {
	var s1 = []int{1, 2, 3, 4, 5, 6}
	s2 := s1[:3]
	s3 := s1[3:]
	fmt.Printf("s1: %T\n", s2)
	fmt.Printf("s2: %v\n", s2)
	fmt.Printf("s3: %v\n", s3)
}

func test4() {
	var s1 = []int{1, 2, 3, 4, 5, 6, 7}
	for i, v := range s1 {
		fmt.Printf("v%v: %v\n", i, v)
	}
}

func test5() {
	// crud
	// add
	var s1 []int
	s1 = append(s1, 100)
	s1 = append(s1, 200)
	fmt.Printf("s1: %v\n", s1)

}

func test6() {
	// crud
	// del 2
	var s1 = []int{1, 2, 3, 4, 5}
	var s2 []int
	s2 = append(s2, s1[:2]...)
	s2 = append(s2, s1[3:]...)
	fmt.Printf("s2: %v\n", s2)
	s1 = append(s1[:2], s1[3:]...)

	fmt.Printf("s1: %v\n", s1)

	//update
	s1[1] = 100
	fmt.Printf("s1: %v\n", s1)

}

func test7() {
	var s1 = []int{1, 2, 3, 4, 5}
	s2 := s1
	s2[2] = 100
	fmt.Printf("s1: %v\n", s1)
	fmt.Printf("s2: %v\n", s2)
	var s3 = make([]int, 8)
	copy(s3, s1)
	s3[2] = 400
	s2[2] = 300
	fmt.Printf("------------------------------")
	fmt.Printf("s1: %v\n", s1)
	fmt.Printf("s2: %v\n", s2)
	fmt.Printf("s3: %v\n", s3)
}
func main() {
	//test1()
	//test2()
	// test3()
	// test4()
	// test5()
	// test6()
	test7()
}
*/