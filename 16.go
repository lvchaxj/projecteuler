package main

import "fmt"

func main() {
	s := []int{2}

	for i := 1; i < 1000; i++ {
		c := 0
		for j := 0; j < len(s); j++ {
			c += 2 * s[j]
			s[j] = c % 10
			c = c / 10
		}
		if c > 0 {
			s = append(s, c)
		}
	}

	n := 0
	for j := 0; j < len(s); j++ {
		n += s[j]
	}

	fmt.Println(n)
}
