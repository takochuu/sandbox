package conflict

import (
	"fmt"
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
