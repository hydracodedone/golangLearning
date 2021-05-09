package package2

import "factoryAutoRegisterDemo/factory"

type animal struct {
	name string
	age  int
}

func initAnimal(name string, age int) interface{} {
	return interface{}(&animal{
		name: name,
		age:  age,
	})
}

func initPeople() func(string, int) interface{} {
	return func(name string, age int) interface{} {
		return &animal{
			name,
			age,
		}
	}
}
func init() {
	factory.GlobalFactory.Register("people", initPeople)
}
