package main

import "fmt"

func main() {
	var a int
	fmt.Scan(&a)
	sum, square := calc(a)
	fmt.Println(sum)
	fmt.Println(square)
}
func calc(a int) (int, int) {
	b := a * 2
	c := a * a
	return b, c
}
