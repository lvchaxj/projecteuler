package main

import "fmt"

func main() {
	var n1, n2 int

	for i := 1; i <= 100; i++ {
		n1 += i * i
		n2 += i
	}

	fmt.Println(n2*n2 - n1)
}
