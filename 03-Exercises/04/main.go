package main

import "fmt"

func main() {
  var myBool bool
  myBool = (true && false) || (false && true) || !(false && false)
  fmt.Println(myBool)
}
