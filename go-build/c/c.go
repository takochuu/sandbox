package c

import "fmt"

type C struct {
}

func init() {
	fmt.Println("C")
}

func NewC() C {
	return C{}
}

func (c *C) Piyo() string {
	return fmt.Sprintf("HogeFugaPiyo")
}
