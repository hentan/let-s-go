package main

import "fmt"

func main() {
	const (
		myConst1 = 1 << iota
		myConst2
		myConst3
		myConst4
		myConst5
	)

	fmt.Println(myConst1, myConst2, myConst3, myConst4, myConst5)
}