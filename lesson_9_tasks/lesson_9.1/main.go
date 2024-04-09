package main

import (
	"errors"
	"fmt"
)

func fruitMarket(fruits_map map[string]int, fruit string) (cnt int, err error) {
	val, ok := fruits_map[fruit]
	if ok {
		return val, nil
	} else {
		err := "не найден фрукт в корзине"
		return 0, errors.New(err)
	}
}

func main(){
	var fruit string
	fmt.Scan(&fruit)
	fruits_map := map[string] int{
		"апельсин" : 5, "яблоки" : 3, "сливы" : 1, "груши" : 0,
	}

	res, err := fruitMarket(fruits_map, fruit)
	if err != nil{
		fmt.Println(err)
	} else {fmt.Println(res)}
}