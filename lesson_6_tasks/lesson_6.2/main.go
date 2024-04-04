package main

import (
	"fmt"
)


func main() {
	type ctr struct {
		ID     int
		Number string
		Date   string
	}

	contract := ctr{
		ID:     1,
		Number: "#000A101\t01", //можно просто добавить обратный слеш перед n
		Date:   "2024-01-31",
	}

	fmt.Printf("%+v\n",contract)
}