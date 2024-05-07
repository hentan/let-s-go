package main

import (
	"encoding/xml"
	"fmt"
)

type contracts struct {
	List []contract `xml:"contract"`
}

type contract struct {
	Number   int    `xml:"number"`
	Landlord string `xml:"landlord"`
	tenat    string `xml:"tenat"`
}

func main() {
	c := contract{
		Number:   1,
		Landlord: "Остап Бендер",
		tenat:    "Шура Балаганов",
	}

	d := contracts{}
	d.List = append(d.List, c)
	res, err := xml.Marshal(d)
	if err != nil {
		fmt.Println("err: ", err)
	}

	fmt.Println(xml.Header, string(res))
}
