package main

import "fmt"

func main() {
	s, c, b := "Go", "Go", "Go"
	fmt.Println(s, &s, c, &c, b, &b)
	/*различные переменные хранятся в разных адресах в памяти*/
}