package main

import "fmt"

type Employee struct {
	Name     string
	Age      int
	Position string
	Salary   int
}

func main() {
	const size = 512
	empls := make([]*Employee, 0, size)

	for {
		fmt.Println("\nВыберите действие:")
		fmt.Println("1 - Добавить нового сотрудника")
		fmt.Println("2 - Удалить сотрудника")
		fmt.Println("3 - Вывести список сотрудников")
		fmt.Println("4 - Выйти из программы")

		var cmd int
		fmt.Print("Введите команду: ")
		fmt.Scanln(&cmd)

		switch cmd {
		case 1:
			if len(empls) >= size {
				fmt.Println("Достигнуто максимальное количество сотрудников")
				continue
			}
			var empl Employee
			fmt.Print("Имя: ")
			fmt.Scanln(&empl.Name)
			fmt.Print("Возраст: ")
			fmt.Scanln(&empl.Age)
			fmt.Print("Позиция: ")
			fmt.Scanln(&empl.Position)
			fmt.Print("Зарплата: ")
			fmt.Scanln(&empl.Salary)
			empls = append(empls, &empl)
			fmt.Println("Сотрудник успешно добавлен")
		case 2:
			if len(empls) == 0 {
				fmt.Println("Список сотрудников пуст")
				continue
			}
			var index int
			fmt.Print("Введите индекс сотрудника для удаления (от 1 до ", len(empls), "): ")
			fmt.Scanln(&index)
			if index < 1 || index > len(empls) {
				fmt.Println("Некорректный индекс")
				continue
			}
			empls = append(empls[:index-1], empls[index:]...)
			fmt.Println("Сотрудник успешно удален")
		case 3:
			if len(empls) == 0 {
				fmt.Println("Список сотрудников пуст")
				continue
			}
			fmt.Println("Список сотрудников:")
			for i, empl := range empls {
				fmt.Printf("%d. Имя: %s, Возраст: %d, Позиция: %s, Зарплата: %d\n",
					i+1, empl.Name, empl.Age, empl.Position, empl.Salary)
			}
		case 4:
			fmt.Println("Выход из программы")
			return
		default:
			fmt.Println("Некорректная команда")
		}
	}
}
