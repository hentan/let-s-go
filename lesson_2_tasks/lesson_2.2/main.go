package main

import "fmt"

func main() {
	a, b := 16, 3
	res := a / b
	rem := a % b
	fmt.Printf("Результат: %d, остаток от деления: %d, тип результата: %T", res, rem, res)

}