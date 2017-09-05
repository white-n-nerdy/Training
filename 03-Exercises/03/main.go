package main

import "fmt"

func main() {
  greatestNumber(1,2,3,4,10,50,6,7,8,9)
}

func greatestNumber(numbers ...int) {
  var x int
  for _, v := range(numbers) {
    if x < v {
      x = v
    }
  }
  fmt.Println(x)
}
