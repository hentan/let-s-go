package main

import "fmt"



func main() {
	target_animals := []string{"слон", "бегемот", "осьминог"}
	count_animals := map[string]int{
		"слон" : 3, "бегемот" : 0, "носорог" : 5, "лев" : 1,
	}
	for _, val := range target_animals{
		an1, ok := count_animals[val]
		if ok{
			fmt.Printf("Животное : %s, количество: %d (есть в карте : %v)\n", val,an1, count_animals)
		}
	}

}