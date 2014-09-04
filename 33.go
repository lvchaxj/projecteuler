package main

import "fmt"

func main() {
	for k := 1; k < 10; k++ {
		for i := 1; i < 10; i++ {
			for j := i + 1; j < 10; j++ {
				a, b := 0, 0
				c := float64(i) / float64(j)
				a, b = 10*k+i, 10*k+j
				if float64(a)/float64(b) == c {
					fmt.Println(i, j)
				}
				a, b = 10*k+i, 10*j+k
				if float64(a)/float64(b) == c {
					fmt.Println(i, j)
				}
				a, b = 10*i+k, 10*k+j
				if float64(a)/float64(b) == c {
					fmt.Println(i, j)
				}
				a, b = 10*i+k, 10*j+k
				if float64(a)/float64(b) == c {
					fmt.Println(i, j)
				}
			}
		}
	}

}
