package main

import (
	"fmt"
)

func main() {
	var number int
	_, _ = fmt.Scan(&number)
	previous := number - 1
	next := number + 1
	fmt.Println("Следующее за числом", number, "число:", next)
	fmt.Println("Для числа", number, "предыдущее число:", previous)
}
