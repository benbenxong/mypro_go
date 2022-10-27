package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type myDogFunc func() interface{}
type myDogStruct struct {
	funcs []myDogFunc
	data  chan interface{}
	wg    *sync.WaitGroup
}

func (this *myDogStruct) Add(f myDogFunc) {
	this.funcs = append(this.funcs, f)
}
func (this *myDogStruct) do() {
	for _, f := range this.funcs {
		this.wg.Add(1)
		go func() {
			defer this.wg.Done()
			this.data <- f()
		}()
	}
}
func (this *myDogStruct) Range(f func(value interface{})) {
	this.do() //执行所有协程
	go func() {
		defer close(this.data) //关闭协程
		this.wg.Wait()
	}()
	for item := range this.data {
		f(item)
	}
}
func Mydog() *myDogStruct {
	return &myDogStruct{data: make(chan interface{}), wg: &sync.WaitGroup{}}
}

func main() {
	dog := Mydog()
	for i := 0; i < 5; i++ {
		dog.Add(func() interface{} {
			time.Sleep(time.Second * time.Duration(rand.Intn(4)))
			return rand.Intn(10) //好比业务数据
		})
	}
	dog.Range(func(value interface{}) {
		fmt.Println(value)
	})
}
