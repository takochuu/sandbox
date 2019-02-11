package main

import (
	"fmt"
)

var limit = make(chan int, 3)
var work []func()

func main() {
	work = append(work, func() {
		fmt.Println("hogehoge")
	})
	work = append(work, func() {
		fmt.Println("fugafuga")
	})
	work = append(work, func() {
		fmt.Println("piyopiyo")
	})
	for _, w := range work {
		go func(w func()) {
			limit <- 1
			w()
			<-limit
		}(w)
	}

	select {
	/*
		case v := <-limit:
			fmt.Println(v)
		default:
			fmt.Println("no value")
	*/
	}
}
