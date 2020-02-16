package main

import (
	"fmt"
)

func main() {
	var a,b int
	a = 3
	b = 2
	/*fmt.Scan(&a, &b) */

	if (a * b)%2 == 0{
	fmt.Println("偶数")
	}else{ 
	fmt.Println("奇数")
	}
}
