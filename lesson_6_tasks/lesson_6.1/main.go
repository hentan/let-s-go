package main

import (
	"fmt"
	"strings"
)

type Contract struct {
	ID     int
	Number string
	Date   string
}

func main() {
	contract := Contract{
		ID:     1,
		Number: "#000A\n101", //можно просто добавить обратный слеш перед n
		Date:   "2024-01-31",
	}

	contract.Number = strings.ReplaceAll(contract.Number, "\n", "\\n")
	fmt.Printf("%+v\n",contract)
}