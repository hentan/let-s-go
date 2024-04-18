package main

import (
	"errors"
	"fmt"
)

var internalError = errors.New("ошибка1")

func getError(level int) error {
    level1Err := fmt.Errorf("ошибка2:%w", internalError)
    if level == 1 {
        return level1Err
    }
    if level == 2 {
        return fmt.Errorf("ошибка3:%w", level1Err)
    }

    return internalError
}

func main() {
	err:= getError(2)

	fmt.Println(err)
}