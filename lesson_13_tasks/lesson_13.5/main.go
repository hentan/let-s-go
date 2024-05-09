package main

import (
	"fmt"
	v10 "v10"
	v11 "v11"
	v20 "v20"
	v21 "v21"
)

func main() {
	err := v10.Do("patients.txt", "res.txt")
	if err != nil {
		fmt.Println("ошибка отработки функции v1.0:", err)
	}

	err = v11.Do("patients.txt", "res2.txt")
	if err != nil {
		fmt.Println("ошибка отработки функции v1.1:", err)
	}

	err = v20.Do("patients.txt", "res3.txt")
	if err != nil {
		fmt.Println("ошибка отработки функции v2.0:", err)
	}

	err = v21.Do("patients.txt", "res4.txt")
	if err != nil {
		fmt.Println("ошибка отработки функции v2.1", err)
	}
}
