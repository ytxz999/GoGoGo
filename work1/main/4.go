package main

import (
	"fmt"
	"math"
)

func main() {
	var N int
	fmt.Scan(&N)
	if isPrime(N) {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

func isPrime(num int) bool {
	if num == 1 {
		return false
	}
	if num == 2 {
		return true
	}
	if num%2 == 0 {
		return false
	}
	for i := 3; float64(i) < math.Sqrt(float64(num)); i = i + 2 {
		if num%i == 0 {
			return false
		}
	}
	return true
}
