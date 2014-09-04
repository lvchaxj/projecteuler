package main

import "fmt"

func main() {
	r := 0
	for i := 1; i <= 1000; i++ {
		v := i
		for j := 1; j < i; j++ {
			v *= i
			v %= 10000000000
		}
		r += v
		r %= 10000000000
	}

	fmt.Println(r)
}
