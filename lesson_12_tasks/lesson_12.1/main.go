package main

import "fmt"

func main() {
	a := 1
	do(a)
}
func do(v any) {
	a := 10

	switch v := v.(type) {
    case int:
        a+=v
	default:
		fmt.Println("была передана переменная отличная от типа int")
		return
	}

	fmt.Println(a)
}