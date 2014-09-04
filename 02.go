package main

import "fmt"

func main() {
	var a, b uint = 1, 1
	var sum uint = 0

	for a < 4000000 {
		a, b = a+b, a
	}

	if a%2 == 0 {
		sum = b
	} else if b%2 == 0 {
		sum = a + b
	} else {
		sum = a
	}

	fmt.Println((sum - 1) / 2)
}
