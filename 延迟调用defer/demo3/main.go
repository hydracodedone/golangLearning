package main

import "fmt"

func demo1() {
	function := func() int {
		res := 1
		defer func() {
			res += 1
		}()
		//RET_VALUE=RES
		//RES+=1
		//ATOMIC RETURN RET_VAVLUE
		return res
	}
	fmt.Printf("the return value is %d\n", function())
}
func demo2() {
	function := func() (result int) {
		defer func() {
			result += 1
		}()
		//RET_VALUE->result
		//result = 1
		//result += 1
		//ATOMIC RETURN RET_VAVLUE->result
		return 0
	}
	fmt.Printf("the return value is %d\n", function())
}
func demo3() {
	function := func() (result int) {
		res := 0
		defer func() {
			res += 1
		}()
		//RET_VALUE->result
		//result=res=0
		//res +=1
		//ATOMIC RETURN RET_VAVLUE->result
		return res
	}
	fmt.Printf("the return value is %d\n", function())
}
func demo4() {
	function := func() (result int) {
		result = 0
		defer func() {
			result += 1
		}()
		//result=0
		//RET_VALUE->result
		//result +=1
		//ATOMIC RETURN RET_VAVLUE->result
		return
	}
	fmt.Printf("the return value is %d\n", function())
}
func main() {
	demo1()
	demo2()
	demo3()
	demo4()
}
