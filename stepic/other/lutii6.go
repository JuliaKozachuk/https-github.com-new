package main

import "fmt"

func main() {
	var a int
	fmt.Scan(&a)
	marsage := mars_age(a)
	fmt.Println(marsage)
}
func mars_age(a int) int {

	var avgPlanet float64
	avgPlanet = 687.00 / 365.00
	fmt.Println(avgPlanet)
	//ap := int(avgPlanet)
	//fmt.Println(ap)
	b := float64(a)

	marstime := b / avgPlanet
	fmt.Println(marstime)
	//fmt.Println(marstime)
	var intmarstime = int(marstime)
	return intmarstime
}
