package main

import (
	"fmt"

	"github.com/takochuu/sandbox/go-build/a"
)

func main() {
	a := a.NewA()
	fmt.Println(a.Hoge())
	return
}
