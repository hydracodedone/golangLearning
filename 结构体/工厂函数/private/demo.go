package demo

type privateStruct struct {
	name string
	age  int
}
type inner struct {
	Age int
}
type PublicatStruct struct {
	Name string
	age  int
	inner
}

func FactoryForStruct(name string, age int) *privateStruct {
	returnValue := new(privateStruct)
	returnValue.name = name
	(*returnValue).age = age
	return returnValue
}
