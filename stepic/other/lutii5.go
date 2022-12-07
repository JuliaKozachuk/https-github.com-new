package main

import "fmt"

func main() {
	var a int
	var c int
	var b string
	fmt.Scan(&a, &c, &b)
	num, str := one_or_two(a, c, b)
	fmt.Println(num, str)
}
func one_or_two(a int, c int, b string) (int, string) {
	var resn int
	var resst string
	if b == "one" {
		resn = a
		resst = b

	} else if b == "two" {
		resn = c
		resst = b
	} else {
		resn = 0
		resst = b
	}
	return resn, resst

}
