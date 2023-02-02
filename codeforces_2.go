package main

import "fmt"

func main() {
	var t, n, p int 
	fmt.Scan(&t) //кол-во наборов входных данных
	for i := 1; i <= t; i++ {
		fmt.Scan(&n)
		// for i:=1;i<=n;i++ {
			
		// }
		var Items []int = make([]int, n)
		fmt.Scan(&p)
		for i:=1;i<=n;i++ {
			Items = append(Items, p)
		}  
	}
fmt.Println(t, n, p)
}