package main

import "fmt"

var c = make(chan int, 10)
var a string

func f() {
	a = "hello, world"
	<-c
}

func main() {
	go f()
	c <- 0
	fmt.Print(a)
}
