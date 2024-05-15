package main

import (
	"fmt"
	"strconv"
	"sync"
)

// Определяем интерфейс Meteo
type Meteo interface {
	ReadTemp() string
	ChangeTemp(v string)
}

type MeteoType struct {
	temp string
	mu   sync.Mutex
}

func (m *MeteoType) ReadTemp() string {
	m.mu.Lock()
	defer m.mu.Unlock()
	return m.temp
}

func (m *MeteoType) ChangeTemp(s string) string {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.temp = s
	return m.temp
}

func main() {
	tempStr := &MeteoType{temp: "20"}
	wg := sync.WaitGroup{}
	wg.Add(200)

	for i := 0; i < 100; i++ {
		go func() {
			defer wg.Done()
			fmt.Println("Текущая температура: ", tempStr.ReadTemp())
		}()

		go func(i int) {
			defer wg.Done()
			fmt.Println("Измененная температура: ", tempStr.ChangeTemp(strconv.Itoa(i)))
		}(i)
	}

	wg.Wait()
}
