package main

import "fmt"

func main() {
	s := []int{1}
	t := []int{1}

	n := 2
	for len(s) < 1000 {
		n++
		ss := make([]int, len(s))
		copy(ss, s)

		c := 0
		for i := 0; i < len(t); i++ {
			c += s[i] + t[i]
			ss[i] = c % 10
			c /= 10
		}
		for i := len(t); i < len(ss); i++ {
			c += s[i]
			ss[i] = c % 10
			c /= 10
		}
		if c > 0 {
			ss = append(ss, c)
		}

		s, t = ss, s
	}

	fmt.Println(n)
}
