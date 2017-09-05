package main

import (
	"fmt"
	"sort"
)

func main() {

	type people []string
	studyGroup := people{"Zeno", "John", "Al", "Jenny"}

	sort.Strings(studyGroup)
  fmt.Printf("%T\n", studyGroup)
	fmt.Println("Sorted: ", studyGroup)

  sort.Sort(sort.Reverse(sort.StringSlice(studyGroup)))
  fmt.Println("Reverse Sorted:", studyGroup)
}
