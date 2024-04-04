package main

import "fmt"

type square int

func String(s square) string {
    return fmt.Sprintf("%d м²", s)
}

func main() {
    var s square = 34
    s += 10

    fmt.Println("Результат:", String(s))
}