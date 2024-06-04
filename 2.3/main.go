package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	_ = scanner.Scan()
	division := scanner.Text()
	_ = scanner.Scan()
	s1 := scanner.Text()
	_ = scanner.Scan()
	s2 := scanner.Text()
	_ = scanner.Scan()
	s3 := scanner.Text()
	fmt.Print(s1, division, s2, division, s3)
}
