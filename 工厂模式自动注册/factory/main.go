package factory

import "fmt"

type structInitFunciton func() func(string, int) interface{}
type factory struct {
	factoryMap map[string]structInitFunciton
}

//GolabalFactory is a golobal
var GolabalFactory *factory = new(factory)

func (f *factory) Register(name string, s structInitFunciton) {
	res := f.factoryMap[name]
	if res == nil {
		f.factoryMap[name] = s
	} else {
		fmt.Println("already registered")
	}
}
func (f *factory) Create(name string) structInitFunciton {
	res := f.factoryMap[name]
	if res != nil {
		return res
	} else {
		fmt.Println("not register yet")
		return nil
	}
}
func init() {
	GolabalFactory.factoryMap = make(map[string]structInitFunciton, 5)
}
