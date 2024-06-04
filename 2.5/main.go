package main

import (
	"fmt"
	"strconv"
)

func main() {
	var number uint16
	_, _ = fmt.Scan(&number)
	n1 := number % 10
	n2 := (number % 100) / 10
	n3 := (number % 1000) / 100
	str1 := strconv.Itoa(int(n1))
	str2 := strconv.Itoa(int(n2))
	str3 := strconv.Itoa(int(n3))
	fmt.Print(str1 + str2 + str3)
}
