// +build prod

package env

type TestStruct struct {
	FirstName string `xorm:"first_name"`
	LastName  string `xorm:"-"`
}

func NewTestStruct() TestStruct {
	return TestStruct{}
}
