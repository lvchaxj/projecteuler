package main

import "fmt"

func main() {
	for a := 1; ; a++ {
		if (1000000-2000*a)%(2000-2*a) == 0 {
			b := (1000000 - 2000*a) / (2000 - 2*a)
			c := 1000 - a - b
			if a*a+b*b == c*c {
				fmt.Println(a * b * c)
				break
			}
		}
	}
}
