package main

import "fmt"
import "math"

func main() {
	prime := []int{2}
	v := 2
	for n := 3; n < 2000000; n += 2 {
		b := true
		sq := int(math.Sqrt(float64(n)))
		for j := 0; j < len(prime) && prime[j] <= sq; j++ {
			if n%prime[j] == 0 {
				b = false
				break
			}
		}
		if b {
			prime = append(prime, n)
			v += n
		}
	}
	fmt.Println(v)
}
