package main

import (
	"fmt"
	"sort"
)

func main() {

  s := []string{"Zeno", "John", "Al", "Jenny"}

	sort.Strings(s)
	fmt.Println("Sorted: ", s)

  sort.Sort(sort.Reverse(sort.StringSlice(s)))
  fmt.Println("Reverse Sorted:", s)
}
