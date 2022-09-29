package main

import "fmt"

//快捷：tyi + tab
type USB interface {
	read()
	write()
}

type Computer struct {
	name string
}

func (c Computer) read() {
	fmt.Printf("c.name: %v\n", c.name)
	fmt.Println("read...")
}

func (c Computer) write() {
	fmt.Printf("c.name: %v\n", c.name)
	fmt.Println("write...")
}

type Pet interface {
	eat()
}

type Cat struct {
}

type Dog struct {
}

func (dog Dog) eat() {
	fmt.Println("dog eat...")
}
func (cat Cat) eat() {
	fmt.Println("cat eat...")
}
func main() {
	// c := Computer{name: "苹果"}
	// c.read()
	// c.write()
	var p Pet
	p = Dog{}
	p.eat()
	p = Cat{}
	p.eat()

}
