package main

import "fmt"

func main() {
	var sum int
	for i := 3; i < 1000; i += 3 {
		sum += i
	}
	for i := 5; i < 1000; i += 5 {
		sum += i
	}
	for i := 15; i < 1000; i += 15 {
		sum -= i
	}
	fmt.Println(sum)
}
