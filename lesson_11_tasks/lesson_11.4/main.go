package main

import (
	"errors"
	"fmt"
)

type myFirstError struct{
    message string
}

func (m *myFirstError) Error() string{
    return m.message
}

func main() {
	err1 := &myFirstError{"ошибка1"}
    err2 := fmt.Errorf("ошибка2:%w", err1)
    err3 := fmt.Errorf("ошибка3:%w", err2)

    myErr := &myFirstError{message: "ошибка1"}
   
    res := errors.As(err3, &myErr)
    if res{
        fmt.Println("ошибка найдена в цепочке ошибок")
        return
    }
    fmt.Println("ошибка не найдена в цепочке ошибок")
}