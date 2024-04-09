package main

import "fmt"

func main() {
	sl := []int{1, 2, 3}
	for _, v := range sl {
		fmt.Println("v1: ", v)
		for _, v2 := range sl {
			if v2 >1{
				continue
			}
			fmt.Println("v2: ", v2)
			for _, v3 := range sl {
				if v3 >1{
					continue
				}
				fmt.Println("v3: ", v3)
				for _, v4 := range sl {
					if v4 >2{ //по условию нужно 1, но в выводе как будто 2?
						continue
					}
					fmt.Println("v4: ", v4)
				}
			}
		}
	}
}