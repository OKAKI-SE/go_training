package main

import (
	"fmt"
)

func main() {
	var n, a, b int
	fmt.Scan(&n, &a, &b)

	var sum int
	var total int

	for i := 1; i <= n; i++{
		for i := 0 {
			sum += i % 10
			i /= 10
		}

		if sum >= a && sum <= b {
			total += i
		}
			fmt.Println()
	}
}
