package main

import "fmt"

func main() {
	sum := 0
	for v := 1; v < 1000000; v += 2 {
		vv := v
		u := 0
		for vv > 0 {
			u = u*10 + vv%10
			vv /= 10
		}
		if u != v {
			continue
		}

		vv = v
		b := []byte{}
		for vv > 0 {
			b = append(b, byte(vv%2))
			vv /= 2
		}
		f := true
		for i := 0; i < len(b); i++ {
			if b[i] != b[len(b)-1-i] {
				f = false
				break
			}
		}

		if f == false {
			continue
		}

		fmt.Println(v)

		sum += v
	}

	fmt.Println(sum)
}
