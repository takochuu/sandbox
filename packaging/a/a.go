package a

import "github.com/takochuu/sandbox/packaging/a/internal"

type A struct {
	internal internal.A
}

func NewA() A {
	return A{
		internal.NewA(),
	}
}
