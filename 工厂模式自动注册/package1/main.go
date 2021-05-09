package package1

import (
	"factoryAutoRegisterDemo/factory"
)

type people struct {
	name string
	age  int
}

func initPeople() func(string, int) interface{} {
	return func(name string, age int) interface{} {
		return &people{
			name,
			age,
		}
	}
}
func init() {
	factory.GlobalFactory.Register("people", initPeople)
}
