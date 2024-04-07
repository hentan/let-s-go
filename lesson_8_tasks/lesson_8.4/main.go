package main

import "fmt"

var Empty struct{}

func main() {
	animals := map[string]struct{}{
	"слон" : Empty,
	"бегемот" : Empty,
	"носорог" : Empty,
	"лев" : Empty,
	}

	animals["выдра"] = Empty

	fmt.Println(animals)
}