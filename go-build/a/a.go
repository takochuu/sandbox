package a

import (
	"fmt"

	"github.com/takochuu/sandbox/go-build/b"
)

type A struct {
}

func init() {
	fmt.Println("A")
}

func NewA() A {
	return A{}
}

func (a *A) Hoge() string {
	b := b.NewB()
	return b.Fuga()
}
