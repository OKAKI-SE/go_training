package main

import (
	"fmt"
)

func main() {
	var a, b int

	fmt.Println("数字を2つ入力してください（例.1 2）")
	fmt.Scan(&a, &b)
	if (a*b)%2 == 0 {
		fmt.Print("Even") //偶数
	} else {
		fmt.Print("Odd") //奇数
	}
}
