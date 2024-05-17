package main

import (
	"context"
	"fmt"
	"sync"
)

func main() {
	ctx := context.Background()

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	wg := sync.WaitGroup{}
	wg.Add(5)
	for i := 0; i < 5; i++ {
		go func(i int) {
			defer wg.Done()
			for {
				select {
				case <-ctx.Done():
					fmt.Println("отмена горутины ", i)
					return
				default:
					fmt.Println("Работа горутины", i)
				}
			}
		}(i)
	}
	wg.Add(1)

	go func() {
		defer wg.Done()
		fmt.Println("стоп!")
		cancel()
	}()

	wg.Wait()
}
