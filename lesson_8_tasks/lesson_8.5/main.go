package main

import "fmt"

var Empty struct{}

func main() {
	count_animals := map[string]int{
		"слон" : 3, "бегемот" : 0, "носорог" : 5, "лев" : 1,
	}

	count_animals["бегемот"] = 2

	fmt.Println(count_animals)
}