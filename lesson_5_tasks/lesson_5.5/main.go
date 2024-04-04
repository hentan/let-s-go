package main

import "fmt"

func main() {
	a := 5
	change(&a)
	fmt.Println(a)

	/*различные переменные хранятся в разных адресах в памяти*/
}

func change(a *int){
	*a = 8
}