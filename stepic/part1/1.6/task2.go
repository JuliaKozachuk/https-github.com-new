package main

import "fmt"

func GetSomeVars() string {
	fmt.Println("GetSomeVars execution")
	return "getSomeVars result"
}
func main() {

	defer fmt.Println("After work")
	defer func() {
		fmt.Println(GetSomeVars())
	}()
	fmt.Println("Some userful work")

}
