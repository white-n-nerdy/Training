package main

import (
  "fmt"
  "sort"
)

func main() {
  n := []int{7, 4, 8, 2, 9, 19, 12, 32, 3}

  sort.Ints(n)
  fmt.Println("Sorted:", n)

  sort.Sort(sort.Reverse(sort.IntSlice(n)))
  fmt.Println("Reverse Sorted:", n)
}
