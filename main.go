package main

import (
	"fmt"
	"log"
)

func main() {
	var n string
	fmt.Println("Добро пожаловать!")
	fmt.Print("Введите данные: ")
	_, err := fmt.Scan(&n)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Вы ввели следующие данные: ", n)
	fmt.Println("Всего Вам доброго!")
}
