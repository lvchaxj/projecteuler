package main

import "fmt"

func main() {
	sum := 1

	for i := 3; i <= 1001; i += 2 {
		sum += 4 * i * i
		sum -= 6 * (i - 1)
	}

	fmt.Println(sum)
}
