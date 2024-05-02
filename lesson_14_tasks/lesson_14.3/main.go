package main

import "fmt"

func main() {
	ch := make(chan string, 2)
	ch <- "Привет"
	ch <- "буферизированный канал"
	fmt.Println(<- ch)
	fmt.Println(<- ch)
}