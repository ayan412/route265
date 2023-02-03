package main

import "fmt"

func main() {
	
	var t, n, p int 
	fmt.Scan(&t) //кол-во наборов входных данных

	var Items []int = make([]int, n) // n кол-во купленных товаров

	for i:=0; i<t;i++ {
		fmt.Scan(&n)
		for i:=0;i<n;i++ {
			fmt.Scan(&p)
			Items = append(Items, p)
		}
		
	}
	sum := 0
	for _, val := range Items {
		sum += val
	}

fmt.Println(t, n, sum)
}