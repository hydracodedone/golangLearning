package main

import "fmt"

/*
递归
*/
func climb_dfs(level int) int {
	if level == 1 || level == 2 {
		return level
	} else {
		return climb_dfs(level-1) + climb_dfs(level-2)
	}
}
func climb(level int) int {
	result := climb_dfs(level)
	return result
}

func main() {
	result := climb(10)
	fmt.Println(result)
}
