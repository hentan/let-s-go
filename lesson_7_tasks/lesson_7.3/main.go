package main

import "fmt"

func main() {
	arr := [4]string{"яблоко", "груша", "помидор", "абрикос"}
	arr[2] = "персик"
	fmt.Println(arr)
}