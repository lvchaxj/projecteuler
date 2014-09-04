package main

import "fmt"
import "math"

func main() {
	for i := 144; ; i++ {
		n := i * (2*i - 1)

		m := 1 + 8*n
		d := int(math.Sqrt(float64(m)))
		if d*d != m || d%2 == 0 {
			continue
		}

		m = 1 + 24*n
		d = int(math.Sqrt(float64(m)))
		if d*d != m || (1+d)%6 != 0 {
			continue
		}

		fmt.Println(n)
		return
	}
}
