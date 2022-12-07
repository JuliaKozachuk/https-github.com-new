package main

import "fmt"

func double(a int) {
	fmt.Println(a * 2)
}

func main() {
	for i := 4; i > 0; i-- {
		defer double(i)
	}
}

// func main() {
// 	var num int
// 	fmt.Scan(&num)
// 	//var statement bool
// 	statement := isEven(num)
// 	fmt.Println(statement)

// }
// func isEven(a int) bool {
// 	if a%2 == 0 {
// 		return true

// 	} else {
// 		return false
// 	}

// }
