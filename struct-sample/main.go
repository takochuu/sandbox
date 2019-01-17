package main

import (
	"reflect"

	"github.com/k0kubun/pp"

	"github.com/takochuu/sandbox/struct-sample/env"
)

func main() {
	s := env.NewTestStruct()
	t := reflect.TypeOf(s)
	f, _ := t.FieldByName("LastName")
	pp.Println(f.Tag)
}
