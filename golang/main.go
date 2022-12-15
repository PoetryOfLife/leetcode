package main

import (
	"fmt"
	"leetcode/find"
)

type Tag struct {
	Value string
}

func main() {
	fmt.Println(find.BsearchRecursion([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 0, 9, 4))
}
