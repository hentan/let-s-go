package main

import (
	"fmt"
	"sync"
)

func main() {
	stop := make(chan struct{})
	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		i := i
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				select {
				case <-stop:
					fmt.Println("стоп горутина: ", i)
					return
				default:
					fmt.Println("сложные вычисления горутины: ", i)
				}
			}
		}()
	}
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("стоп!")
		close(stop)
	}()
	wg.Wait()
}
