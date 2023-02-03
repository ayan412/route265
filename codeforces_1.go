package main

import "fmt"

func main() {
	var t int 
	var a, b, c float32
	fmt.Scan(&t)
	for i := 1; i <= t; i++ {
		fmt.Scan(&a, &b)
		if (a >= -1000 && a <= 1000) && (b >= -1000 && b <= 1000) {
			c = a + b
			fmt.Println(c)
		}
	}
}
