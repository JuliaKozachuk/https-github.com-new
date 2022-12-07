package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	sum := double_m(a, b)
	fmt.Println(sum)
}
func double_m(a int, b int) int {
	sum := 0
	for i := a; i <= b; i++ {
		//res = a * i
		//i++
		sum += i * i
		fmt.Println(sum)

	}
	return sum

}
