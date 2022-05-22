package main

import "math"

func divide(dividend int, divisor int) int {
	sign, res := -1, 0
	if dividend == 0 {
		return 0
	}

	if divisor == 1 {
		return dividend
	}

	if dividend == math.MinInt32 && divisor == -1 {
		return math.MaxInt32
	}

	if dividend > 0 && divisor > 0 || dividend < 0 && divisor < 0 {
		sign = 1
	}

	dvd, dvs := abs(dividend), abs(divisor)

	for dvd >= dvs {
		temp := dvs
		m := 1
		for temp<<1 <= dvd {
			temp <<= 1
			m <<= 1
		}
		dvd -= temp
		res += m
	}
	return sign * res
}

func abs(a int) int {
	if a >= 0 {
		return a
	} else {
		return a * -1
	}
}
