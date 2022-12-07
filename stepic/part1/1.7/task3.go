package main
import "fmt"
type Person struct {
	ID int
	Name string
	Address string
}
type Account struct {
	ID int
	Name string
	Cleaner func(string)string
	Owner Person
}
func main (){
	var acc = Account{
		ID: 1,
		Name : "rvasily",

	}
	fmt.Printf("%#v\n" acc)
	acc.Owner = Person{2,"Romanov Vasily", "Moskow"}
	fmt.Printf("%#v\n" acc)
}