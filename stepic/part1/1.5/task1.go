//package main

//import "fmt"

//func main() {
//var user map[string]string = map[string]string{
//	"name":     "Vasily",
//	"LastName": "Romanov",
//	}
//	profile := make(map[string]string, 10)
//	mapLegth := len(user)
//	fmt.Println(mapLegth, profile)
//	mName := user["middle name"]
//	fmt.Println("mName:", mName)
//	mName, mNameExist := user["middle name"]
//	fmt.Println("m name:", mName, "m NameExist", mNameExist)
//	_, mNameExist2 := user["middle name"]
//	fmt.Println("middle name:", mNameExist2)
//	delete(user, "LastName")
//	fmt.Println("%#v/n", user)

//var buf0 []iPrintln()
//buf1 := []int{0, 1, 2, 3, 4}
//fmt.Println(buf1[1:4])
//buf2 := []int{42}
//buf3 := make([]int, 0)
//buf4 := make([]int, 5)
//buf5 := make([]int, 5, 10)
//fmt.Println(buf0, buf1, buf2, buf3, buf4, buf5)
//buf0 = append(buf0, 9, 10)
//buf0 = append(buf0, 12)
//buf0 = append(buf0, buf1...)
//fmt.Println(buf0, buf1)
//var buf0Len, buf0Cap int = len(buf0), cap(buf0)
//fmt.Println(buf0Len, buf0Cap)

// }
package main

func singleIn(in int) int {
	return in

}
func multiIn(a, b int, c int) int {
	return a + b + c
}
func nameReturn() (out int) {
	out = 2
	return 3
}
func multipleReturn (in int) (int, error){
	if in>2 {
	return 0 fmt.Errorf ("some error happend")}
}
