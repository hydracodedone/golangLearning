package main

import "fmt"

/*
递归
*/
func getCache(level int, levelSolution []int) (int, bool) {
	if levelSolution[level] == 0 {
		return 0, false
	} else {
		return levelSolution[level], true
	}
}
func setCache(level int, levelSolution []int, value int) {
	levelSolution[level] = value
}
func climb_dfs(level int, levelSolution []int) int {
	if level == 1 || level == 2 {
		_, ok := getCache(level, levelSolution)
		if !ok {
			setCache(level, levelSolution, level)
		}
		return level
	} else {
		finalValue := 0
		for eachLevel := level - 2; eachLevel < level; eachLevel++ {
			value, ok := getCache(eachLevel, levelSolution)
			if !ok {
				value = climb_dfs(eachLevel, levelSolution)
				setCache(eachLevel, levelSolution, value)
			}
			finalValue += value
		}
		return finalValue
	}
}
func climb(level int) int {
	var levelSolution []int = make([]int, level)
	result := climb_dfs(level, levelSolution)
	fmt.Println(levelSolution)
	return result
}

func main() {
	result := climb(10)
	fmt.Println(result)
}
