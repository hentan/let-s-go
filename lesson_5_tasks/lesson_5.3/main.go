package main

import "fmt"

func main() {
	var s string = "Go"
	p:= &s
	*p = "Python"

	fmt.Println(s)
}