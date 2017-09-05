package main

import "fmt"

func main() {

  var smallNum, largeNum int

  fmt.Println("Please enter a small number and a large number")
  fmt.Scan(&smallNum,&largeNum)
  remainder := largeNum % smallNum
  fmt.Printf("%v", remainder)
}
