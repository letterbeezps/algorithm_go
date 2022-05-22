package main

import "fmt"

func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	} else if n == 1 {
		return x
	} else if n < 0 {
		return 1 / myPow(x, -1*n)
	}

	var res float64 = 1
	for n > 0 {
		if n&1 == 1 {
			res *= x
		}
		x *= x
		n /= 2
	}

	return res
}

func main() {
	var x float64 = 2.0000
	var n int = -2

	fmt.Println(myPow(x, n))
}
