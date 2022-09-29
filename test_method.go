/* package main

import "fmt"

type Person struct {
	name string
}

func (recv Person) eat() {
	fmt.Printf("%v,eating...\n", recv.name)
}

func (recv Person) sleep() {
	fmt.Printf("%v,sleeping...\n", recv.name)
}

type Customer struct {
	name string
}

func (recv Customer) login(name string, pwd string) bool {
	fmt.Printf("recv.name: %v\n", recv.name)
	if name == "tom" && pwd == "123" {
		return true
	} else {
		return false
	}
}

func showPerson1(per Person) {
	per.name = "tom..."
}

func showPerson2(per *Person) {
	//自动解引用
	//(*per).name
	per.name = "tom..."
}

func (per Person) showPerson3() {
	per.name = "tom..."
}

func (per *Person) showPerson4() {
	//自动解引用
	//(*per).name
	per.name = "tom..."
}
func main() {
	// per := Person{name: "tom"}
	// per.eat()
	// per.sleep()

	// c := Customer{name: "tom"}
	// fmt.Printf("c.login(\"tom\", \"123\"): %v\n", c.login("tom", "123"))
	p1 := Person{name: "tom"}
	p2 := &Person{name: "tom"}
	fmt.Printf("p1: %T\n", p1)
	fmt.Printf("p2: %T\n", p2)
	fmt.Println("--------------------------")
	showPerson1(p1)
	showPerson2(p2)
	fmt.Printf("p1.name: %v\n", p1.name)
	fmt.Printf("p2.name: %v\n", p2.name)

	fmt.Println("--------------------------")
	p1 = Person{name: "tom"}
	p2 = &Person{name: "tom"}
	p1.showPerson3()
	p2.showPerson4()
	fmt.Printf("p1.name: %v\n", p1.name)
	fmt.Printf("p2.name: %v\n", p2.name)

}
*/