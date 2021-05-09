package pk1

import (
	_ "demo_for_init/pk5"
	"fmt"
)

func init() {
	fmt.Println("this is pk1 begin")
}

var GlobalParam = getGlobalParam()

func getGlobalParam() int {
	fmt.Println("this is pk1 getGlobalParam")
	return 1
}

func init() {
	fmt.Println("this is pk1 end")
}
