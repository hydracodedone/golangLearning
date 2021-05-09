package main

import "fmt"

func funcReturnInterface(name interface{}, age interface{}) interface{} {
	transformName, ok := name.(string)
	transformAge, ok2 := age.(int)
	if ok && ok2 {
		return &struct {
			name string
			age  int
		}{
			name: transformName,
			age:  transformAge,
		}
	} else {
		return nil
	}
}

func demoForFunctionReturnInterface() {
	res := funcReturnInterface("Hydra", 23)
	transFormRes, ok := res.(*struct {
		name string
		age  int
	})
	if ok {
		fmt.Printf("The transformRes type is %T\n", transFormRes)
	}
	fmt.Printf("The res type is %T\n", res)

}

func demoForTransformSpecificTypeToInterface() interface{} {
	age := 23
	return interface{}(age)
}

func main() {
	demoForFunctionReturnInterface()
	res := demoForTransformSpecificTypeToInterface()
	fmt.Printf("The type is %T\n", res)
}
