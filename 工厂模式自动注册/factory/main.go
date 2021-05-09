package factory

import "fmt"

type structInitFunction func() func(string, int) interface{}
type factory struct {
	factoryMap map[string]structInitFunction
}

//GlobalFactory is a global
var GlobalFactory *factory = new(factory)

func (f *factory) Register(name string, s structInitFunction) {
	res := f.factoryMap[name]
	if res == nil {
		f.factoryMap[name] = s
	} else {
		fmt.Println("already registered")
	}
}
func (f *factory) Create(name string) structInitFunction {
	res := f.factoryMap[name]
	if res != nil {
		return res
	} else {
		fmt.Println("not register yet")
		return nil
	}
}
func init() {
	GlobalFactory.factoryMap = make(map[string]structInitFunction, 5)
}
