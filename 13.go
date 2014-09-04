package main

import "fmt"

func main() {
	var s string
	fmt.Scanf("%s", &s)

	sum := make([]int, len(s))
	for i := 0; i < len(s); i++ {
		sum[i] = int(s[len(s)-1-i] - '0')
	}

	for j := 1; j < 100; j++ {
		fmt.Scanf("%s", &s)
		v := 0
		for i := 0; i < len(s); i++ {
			v += int(s[len(s)-1-i] - '0')
			v += sum[i]
			sum[i] = v % 10
			v /= 10
		}
		for i := len(s); i < len(sum); i++ {
			v += sum[i]
			sum[i] = v % 10
			v /= 10
		}
		if v > 0 {
			sum = append(sum, v)
		}
	}

	for i := 1; i <= 10; i++ {
		fmt.Printf("%d", sum[len(sum)-i])
	}
	fmt.Println()
}
