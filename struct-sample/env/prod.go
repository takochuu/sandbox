// +build jp

package env

type TestStruct struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"-"`
}

func NewTestStruct() TestStruct {
	return TestStruct{}
}
