package main

import (
	"errors"
	"fmt"
)


func main() {
	err1 := errors.New("ошибка1")
    err2 := fmt.Errorf("ошибка2:%w", err1)
    err3 := fmt.Errorf("ошибка3%w", err2)
    res := errors.Is(err3, err1)
    if res {
        fmt.Println("ошибка1 была среди проблемных кейсов")
    }
}