package main

import "fmt"

func main() {
	fmt.Println("Начало главной функции")
	doSomething()
	fmt.Println("Конец главной функции")
}

func doSomething() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Паника обнаружена:", err)
		}
	}()

	fmt.Println("Выполнение функции doSomething")
	panic("Что-то пошло не так!")
	fmt.Println("Этот код не выполнится из-за паники")
}
