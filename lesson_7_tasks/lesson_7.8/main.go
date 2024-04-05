package main

import (
	"fmt"
	"sort"
)

func main() {
	num := 4
	sl := []int{1, 2, 3, 4, 5, 6}
	a := sort.IntSlice(sl[0:])
	index := sort.SearchInts(a, num)
	sl = append(sl[:index], sl[index+1:]...)
	fmt.Println(sl)
}