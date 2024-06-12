package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 2) // Буферизованный канал с размером буфера 2

	// Запуск нескольких горутин для записи в канал
	go func() {
		for i := 0; i < 5; i++ {
			ch <- i
			fmt.Println("Записано:", i)
		}
		close(ch)
	}()

	// Чтение из канала с таймаутом
	timeout := time.After(3 * time.Second)
	for {
		select {
		case val, ok := <-ch:
			if !ok {
				fmt.Println("Канал закрыт")
				return
			}
			fmt.Println("Прочитано:", val)
		case <-timeout:
			fmt.Println("Тайм-аут чтения")
			return
		}
	}
}
