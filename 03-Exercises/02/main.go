package main

import "fmt"

func main() {
	var varNumber int
	fmt.Println("Please enter a number")
	fmt.Scan(&varNumber)

	myFunc := func(varNumber int) (int, bool) {
		var i int
		var j bool

		i = varNumber / 2
		j = varNumber%2 == 0 //Evanluates to true or false

		return i, j
	}

	fmt.Println(myFunc(varNumber))
}
