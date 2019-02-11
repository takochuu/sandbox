package main

import "fmt"

var a string

func f() {
	fmt.Println(a)
}

func main() {
	a = "hello, world"
	go f()
}
