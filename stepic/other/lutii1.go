package main

import "fmt"

func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	x := max(a, b)
	fmt.Println(x)
}
