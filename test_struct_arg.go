package main

import "fmt"

type Person struct {
	id   int
	name string
}

func showPerson(per Person) {
	per.id = 101
	per.name = "kite"
	fmt.Printf("per: %v\n", per)
}

func showPerson2(per *Person) {
	per.id = 102
	per.name = "kite"
	fmt.Printf("per: %v\n", per)
}

func main() {
	var tom1 = new(Person)
	tom1.id = 101
	tom1.name = "tom1"
	fmt.Printf("tom1: %T\n", tom1)
	tom := Person{
		id:   101,
		name: "tom",
	}
	fmt.Printf("tom: %v\n", tom)
	fmt.Println("------------------")
	showPerson(tom)

	fmt.Printf("tom: %v\n", tom)
	fmt.Println("-------------------")
	fmt.Printf("tom1: %v\n", tom1)
	showPerson2(tom1)
	fmt.Printf("tom1: %v\n", tom1)

}
