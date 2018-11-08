// +build wireinject
package _interface

import "github.com/google/go-cloud/wire"

type PiyoI interface {
	Piyo() string
}

type Hoge string

func (b *Hoge) Piyo() string {
	return string(*b)
}

func ProvideHoge() *Hoge {
	b := new(Hoge)
	*b = "Hello, World!"
	return b
}

var BarFooer = wire.NewSet(ProvideHoge, wire.Bind(new(PiyoI), new(Hoge)))
