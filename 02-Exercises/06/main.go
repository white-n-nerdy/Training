package main

import "fmt"

func main() {
	for i := 0; i <= 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Printf("%s\n", "FizzBuzz")
		} else if i%5 == 0 {
			fmt.Printf("%s\n", "Buzz")
		} else if i%3 == 0 {
			fmt.Printf("%s\n", "Fizz")
		} else {
			fmt.Printf("%d\n", i)
		}
	}
}
