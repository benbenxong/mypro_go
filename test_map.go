/* package main

import "fmt"

func test1() {
	//var m1 map[string]string
	m1 := make(map[string]string)
	m1["name"] = "tom"
	m1["age"] = "20"
	m1["emal"] = "hj@1234.com"
	fmt.Printf("m1: %v\n", m1)

	//get value from key "name"
	key := "name"
	value := m1[key]
	fmt.Printf("value: %v\n", value)
}

func test2() {
	var m1 = map[string]string{"age": "20", "emal": "hj@1234.com", "name": "tom"}
	fmt.Printf("m1: %v\n", m1)
	var k1 = "name"
	var k2 = "age1"
	v, ok := m1[k1]
	fmt.Printf("v: %v\n ok: %v\n", v, ok)
	v, ok = m1[k2]
	fmt.Printf("v: %v\n ok: %v\n", v, ok)
}

func test3() {
	var m1 = map[string]string{"age": "20", "emal": "hj@1234.com", "name": "tom"}
	for key := range m1 {
		fmt.Printf("key: %v\n", key)

	}
	fmt.Printf("\"\": %v\n", "-----------------")
	for key, v := range m1 {
		fmt.Printf("k:%v v: %v\n", key, v)

	}
	fmt.Printf("modified!")
}
func main() {
	// test1()
	// test2()
	test3()
}
*/