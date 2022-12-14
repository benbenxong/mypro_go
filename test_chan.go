/* package main

import (
	"fmt"
	"math/rand"
	"time"
)

var values = make(chan int)

func send() {
	rand.Seed(time.Now().UnixNano())
	value := rand.Intn(10)
	fmt.Printf("send value: %v\n", value)
	time.Sleep(time.Second * 3)
	values <- value
}

func main() {
	defer close(values)
	go send()
	fmt.Println("wait...")
	value := <-values
	fmt.Printf("receive value: %v\n", value)
	fmt.Println("end...")
}
*/