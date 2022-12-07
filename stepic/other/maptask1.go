package main

import "fmt"

func main() {
	var num int
	fmt.Scan(&num)
	sum := 1
	res := 1

	for sum <= num {

		res *= sum
		//fmt.Println(res)
		sum++

	}
	fmt.Println(res)
	//fmt.Println(res * sum)
}
