package main

import (
	"fmt"
)

func main() {
	var num1 float64
	_, _ = fmt.Scan(&num1)
	num2 := int(num1)
	num1 -= float64(num2)
	fmt.Println(num1)
}
