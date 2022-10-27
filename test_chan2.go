package main

import (
	"sync"
	"time"
)

var wg = sync.WaitGroup{}

func pub_job() {
	defer wg.Done()
	time.Sleep(time.Second)
}
func pri_job() {
	defer wg.Done()
	time.Sleep(time.Second * 2)
}
func main() {
	start := time.Now()
	wg.Add(2)
	go pub_job()
	go pri_job()
	wg.Wait()
	println(time.Since(start).Milliseconds() / 1000)
}
