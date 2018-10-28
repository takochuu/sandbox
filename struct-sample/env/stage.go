// +build !prod

package env

type TestStruct struct {
	FirstName string `xorm:"first_name"`
	LastName  string `xorm:"last_name"`
}

func NewTestStruct() TestStruct {
	return TestStruct{}
}
