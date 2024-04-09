package main

import "fmt"

func checkFood(food string) {
	switch food {

	case "апельсин","груша", "яблоко":
		fmt.Println("фрукт")
	
	case "тыква", "огурец", "помидор":
		fmt.Println("овощ")
	default:
		fmt.Println("что-то странное…")
	}
}

func main() {
	food := "огурец"

	checkFood(food)
}