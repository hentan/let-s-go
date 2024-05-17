package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	ctx := context.Background()
	d := time.Now().Add(2 * time.Second)
	ctx, cancel := context.WithDeadline(ctx, d)
	defer cancel()
	wg := sync.WaitGroup{}
	wg.Add(5)
	for i := 0; i < 5; i++ {
		go func(i int) {
			defer wg.Done()
			for {
				select {
				case <-ctx.Done():
					fmt.Println("отмена горутины ", i, ctx.Err())
					return
				default:
					fmt.Println("Работа горутины", i)
					time.Sleep(time.Second)
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
