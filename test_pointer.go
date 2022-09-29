/* package main

import "fmt"

func main() {
	var ip *int
	var i = 100
	ip = &i
	fmt.Printf("ip: %v\n", ip)
	fmt.Printf("*ip: %v\n", *ip)

	//因为go指针不能计算，所以指针不能指向数组第一个元素，只能一个个指定
	//var sp [3]*int . 即3个分别指向每个元素的指针。
}
*/