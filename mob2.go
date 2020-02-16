package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	var s string
	fmt.Scan(&s)
	slice := strings.Split(s, "")
	sum := 0
	for _, num := range slice {
		i, _ := strconv.Atoi(num)
		sum += i
	}
	fmt.Println(sum)
}
