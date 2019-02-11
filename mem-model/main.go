package main

import (
	"fmt"
)

var limit = make(chan int, 3)
var work []func()

func main() {
	work = append(work, func() {
		fmt.Println("hogehoge")
		limit <- 1
	})
	work = append(work, func() {
		fmt.Println("fugafuga")
		limit <- 2
	})
	work = append(work, func() {
		fmt.Println("piyopiyo")
		limit <- 3
	})
	for _, w := range work {
		go func(w func()) {
			w()
		}(w)
	}
	fmt.Println(<-limit)
}
