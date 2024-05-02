package main

import "fmt"

func main() {
	ch := make(chan string)
	go func() {
		txt := "Привет, строковый канал!"
		ch <- txt
	}()

	fmt.Println(<-ch)

}