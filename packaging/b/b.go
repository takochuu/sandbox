package b

import "github.com/takochuu/sandbox/packaging/a/internal"

type B struct {
	internal internal.A
}

func NewB() B {
	return B{
		internal: internal.NewA(),
	}
}
