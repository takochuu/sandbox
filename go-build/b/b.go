package b

import (
	"fmt"

	"github.com/takochuu/sandbox/go-build/c"
)

type B struct {
}

func init() {
	fmt.Println("B")
}

func NewB() B {
	return B{}
}

func (b *B) Fuga() string {
	c := c.NewC()
	return c.Piyo()
}
