package main

import (
	"reflect"

	"github.com/acidlemon/go-dumper"
	"github.com/takochuu/sandbox/struct-sample/env"
)

func main() {
	s := env.NewTestStruct()
	t := reflect.TypeOf(s)
	f, b := t.FieldByName("LastName")
	dump.Dump(f)
	dump.Dump(b)
}
