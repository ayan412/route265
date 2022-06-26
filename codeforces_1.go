package main

import "fmt"

func main() {
	var t int
	var a, b, c float32
	fmt.Scan(&t)
	for i := 1; i <= t; i++ {
		fmt.Scan(&a, &b)
		c = a + b
		fmt.Println(c)
	}
}
