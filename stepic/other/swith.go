package main

import "fmt"

var x = 13

func change() {
	x += 1
	fmt.Println(x)
}
func set(x int) {
	x += 3
	fmt.Println(x)
}
func main() {
	change()
	set(x)
	fmt.Println(x)
}

// package main

// import "fmt"

// func main() {
// 	var onenum int

// 	for i := 0; i < 3; i++ {
// 		fmt.Scan(&onenum)

// 		switch onenum {
// 		case 10:
// 			fmt.Println("Десять")
// 		case 1:
// 			fmt.Println("Один")
// 		case 2:
// 			fmt.Println("Два")
// 		case 3:
// 			fmt.Println("Три")
// 		case 4:
// 			fmt.Println("Четыре")
// 		case 5:
// 			fmt.Println("Пять")
// 		case 6:
// 			fmt.Println("Шесть")
// 		case 7:
// 			fmt.Println("Семь")
// 		case 8:
// 			fmt.Println("Восемь")
// 		case 9:
// 			fmt.Println("Девять")
// 		default:
// 			fmt.Println("Ноль")

// 		}

// 	}
// }

// var twonum int
// fmt.Scan(&twonum)

// switch twonum {
// case 10:
// 	fmt.Println("Десять")
// case 1:
// 	fmt.Println("Один")
// case 2:
// 	fmt.Println("Два")
// case 3:
// 	fmt.Println("Три")
// case 4:
// 	fmt.Println("Четыре")
// case 5:
// 	fmt.Println("Пять")
// case 6:
// 	fmt.Println("Шесть")
// case 7:
// 	fmt.Println("Семь")
// case 8:
// 	fmt.Println("Восемь")
// case 9:
// 	fmt.Println("Девять")
// default:
// 	fmt.Println("Ноль")

// }
// var threenum int
// fmt.Scan(&threenum)

// switch threenum {
// case 10:
// 	fmt.Println("Десять")
// case 1:
// 	fmt.Println("Один")
// case 2:
// 	fmt.Println("Два")
// case 3:
// 	fmt.Println("Три")
// case 4:
// 	fmt.Println("Четыре")
// case 5:
// 	fmt.Println("Пять")
// case 6:
// 	fmt.Println("Шесть")
// case 7:
// 	fmt.Println("Семь")
// case 8:
// 	fmt.Println("Восемь")
// case 9:
// 	fmt.Println("Девять")
// default:
// 	fmt.Println("Ноль")

// }
