package main

import "fmt"

/*
回溯来穷举
*/
func handle_func(choices []int, level, state int, pResult *int) {
	if state == level {
		*pResult = *pResult + 1
	}
	for _, step := range choices {
		if step+state > level {
			continue
		}
		handle_func(choices, level, state+step, pResult)
	}
}
func climb(level int) int {
	choices := []int{1, 2}
	state := 0
	result := 0
	var pResult *int = &result
	handle_func(choices, level, state, pResult)
	return result
}

func main() {
	result := climb(10)
	fmt.Println(result)
}
