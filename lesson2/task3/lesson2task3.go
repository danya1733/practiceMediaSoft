package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rows := 4
	cols := 5

	// Создаем двумерный массив заданного размера
	arr := make([][]int, rows)
	for i := range arr {
		arr[i] = make([]int, cols)
	}

	// Создаем множество для хранения уникальных чисел
	uniqueNumbers := make(map[int]bool)

	// Инициализируем генератор случайных чисел
	rand.Seed(time.Now().UnixNano())

	// Заполняем массив случайными уникальными числами
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			// Генерируем случайное число до тех пор, пока оно не будет уникальным
			for {
				num := rand.Intn(rows * cols)
				if !uniqueNumbers[num] {
					arr[i][j] = num
					uniqueNumbers[num] = true
					break
				}
			}
		}
	}

	// Выводим заполненный массив
	for _, row := range arr {
		fmt.Println(row)
	}
}
