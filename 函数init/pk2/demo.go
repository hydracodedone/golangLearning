package pk2

import (
	"demo_for_init/pk4"
	_ "demo_for_init/pk5"
	"fmt"
)

func init() {
	fmt.Println("this is pk2 begin")
	pk4.Demo4()
}
func init() {
	fmt.Println("this is pk2 end")
}
