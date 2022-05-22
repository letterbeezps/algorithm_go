package main

import "fmt"

func addBinary(a string, b string) string {
	n := max(len(a), len(b))
	res := make([]byte, n+1)
	carry := byte(0)

	i, j, k := len(a)-1, len(b)-1, n
	for i >= 0 || j >= 0 {

		cura, curb := byte(0), byte(0)
		if i >= 0 {
			cura = a[i] - 48
		}
		if j >= 0 {
			curb = b[j] - 48
		}
		ret := cura + curb + carry
		carry = ret / 2
		ret %= 2

		res[k] = ret + 48

		i--
		j--
		k--
	}

	if carry == byte(1) {
		res[k] = byte(1) + 48
	} else {
		res = res[1:]
	}

	return string(res)
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func main() {
	a, b := "1010", "1011"
	fmt.Println(addBinary(a, b))
}
