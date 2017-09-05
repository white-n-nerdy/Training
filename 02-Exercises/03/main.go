package main

import "fmt"

func main() {

	var readValue string

	fmt.Print("Please enter your name: ")
	fmt.Scan(&readValue)
	fmt.Println("Hello", readValue)

}
