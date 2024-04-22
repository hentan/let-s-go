package main

import "fmt"

type Formatter interface {
	Format()
}

type Xml struct{}

func (x Xml) Format() {
	fmt.Println("Данные в формате xml")
}

type Csv struct{}

func (c Csv) Format() {
	fmt.Println("Данные в формате csv")
}
func main() {
	x := Xml{}
	Report(x)
	c := Csv{}
	Report(c)
}

func Report(f Formatter) {
	if xml, ok := f.(Xml); ok {
		xml.Format()
	}
	if csv, ok := f.(Csv); ok {
		csv.Format()
	}
}
