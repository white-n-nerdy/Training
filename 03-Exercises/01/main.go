package main

import "fmt"

func main() {
	var varNumber int
	fmt.Println("Please enter a number")
	fmt.Scan(&varNumber)

	firstArg, secondArg := myFunc(varNumber)
	fmt.Println("Half of", varNumber, ":", firstArg, "Is it even?:", secondArg)
}

func myFunc(myNum int) (int, bool) {
	var i int
	var j bool

	i = myNum / 2
	j = myNum%2 == 0

	return i, j
}
