package conflict

import (
	"fmt"
	"sync"
	"time"
)

func Conflict() {
	var data int
	go func() {
		data++
	}()
	time.Sleep(1 * time.Second)
	if data == 0 {
		fmt.Printf("the value is %v", data)
	}
}

func Conflict1() {
	var data int
	go func() {
		data++
	}()
	if data == 0 {
		fmt.Println("the value is 0")
	} else {
		fmt.Printf("the value is %v", data)
	}
}

func Conflict2() {
	var m sync.Mutex
	var value int
	go func() {
		m.Lock()
		value++
		m.Unlock()
	}()
	m.Lock()
	if value == 0 {
		fmt.Println("the value is 0")
	} else {
		fmt.Printf("the value is %v", value)
	}
	m.Unlock()
}
