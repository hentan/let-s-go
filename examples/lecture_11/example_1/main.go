package main

import (
	"errors"
	"log"
	"os"
	"strconv"
)

type User struct {
	Name string
	Age  int
}

func (u User) List(limit, offset int) ([]User, error) {
	if limit < 0 || offset < 0 {
		return []User{}, errors.New("что-то пошло не так")
	}
	if limit == 0 {
		return []User{}, nil
	}
	result := make([]User, 0, limit)
	result = append(result, User{
		Name: "Остап Бендер",
		Age:  44,
	})
	return result, nil
	}

func main(){
	if len(os.Args) != 3 {
		log.Fatalln("Недостаточно данных: limit, offset")}
		
		limit, err := strconv.Atoi(os.Args[1])

		if err != nil {
			log.Fatalln(err)
		}
		offset, err := strconv.Atoi(os.Args[2])
		if err != nil {
			log.Fatalln(err)
		}

		u := User{}
		res, err := u.List(limit, offset)
		if err != nil {
			log.Fatalln("Ошибка:", err)
		}
		log.Printf("Список пользователей: %v", res)
		
}