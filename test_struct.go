/* package main

import "fmt"

func test2() {
	type Person struct {
		id   int
		name string
		age  int
	}

	tom := Person{
		id:   101,
		name: "tom",
		age:  20,
	}

	var p_Person *Person
	p_Person = &tom

	fmt.Printf("tom: %v\n", tom)
	fmt.Printf("p_Person: %p\n", p_Person)
	fmt.Printf("p_Person: %v\n", p_Person)
	fmt.Printf("p_Person: %v\n", *p_Person)
}

func test3() {
	type Person struct {
		id   int
		name string
		age  int
	}

	tom := new(Person)

	fmt.Printf("tom: %T\n", tom)

	tom.id = 101
	(*tom).name = "tom"
	fmt.Printf("tom: %v\n", tom)
	// var p_Person *Person
	// p_Person = &tom

	// fmt.Printf("tom: %v\n", tom)
	// fmt.Printf("p_Person: %p\n", p_Person)
	// fmt.Printf("p_Person: %v\n", p_Person)
	// fmt.Printf("p_Person: %v\n", *p_Person)
}
func main() {
	// test2()
	test3()
}
*/