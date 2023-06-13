package main

import (
	"fmt"
	"log"
)

func main() {
	var n string
	fmt.Print("Введите данные: ")
	_, err := fmt.Scan(&n)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Вы ввели следующие данные: ", n)
}
