package main

import (
	"context"
	"fmt"
)

type Values struct {
	m map[string]string
}

func (v Values) Get(key string) string {
	return v.m[key]
}

type ctxKey int

func main() {
	v := Values{map[string]string{
		"some key1": "some value1",
		"some key2": "some value2",
	}}
	c := context.Background()
	ctx := context.WithValue(c, "val", v)

	fmt.Println(ctx.Value("val").(Values).Get("some key1"))
	fmt.Println(ctx.Value("val").(Values).Get("some key2"))
}

func do(ctx context.Context) {
	var key1, key2 ctxKey = 1, 2
	fmt.Println("some key1: ", ctx.Value(key1))
	fmt.Println("some key2: ", ctx.Value(key2))
}
