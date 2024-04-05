package main

import "fmt"

func main() {
	sl1 := []int{1,2,3}
	sl2 := []int{4,5,6}
	sl := append(sl1, sl2...)
	fmt.Println(sl)
}