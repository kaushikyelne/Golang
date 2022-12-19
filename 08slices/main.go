package main

import (
	"fmt"
	"sort"
)

func main() {
	var scores = []int{}
	fmt.Println(scores)
	scores = append(scores, 8, 10, 14, 18, 20, 1, 5)
	fmt.Println(scores[1:])
	fmt.Println(scores[1:3])
	fmt.Println(sort.IntsAreSorted(scores))
	sort.Ints(scores)
	fmt.Println(scores)
	fmt.Println(scores[:3])

}
