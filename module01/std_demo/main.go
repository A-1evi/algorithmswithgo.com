package main

import (
	"algorithmswithgo/module01"
	"fmt"
)

func main() {
	var n int
	fmt.Scanf("%d", &n)
	for i := 0; i <= n; i++ {
		var a, b int
		fmt.Scanf("%d %d", &a, &b)
		gcd := module01.GCD(a, b)
		fmt.Printf("The gcd of a , b : %d\n", gcd)
	}

}
