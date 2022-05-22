package main

import "fmt"

func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	res := make([]int, len(num1)+len(num2))

	byte1, byte2 := []byte(num1), []byte(num2)
	for i := 0; i < len(byte1); i++ {
		for j := 0; j < len(byte2); j++ {
			res[i+j+1] += int(byte1[i]-'0') * int(byte2[j]-'0')
		}
	}

	for i := len(res) - 1; i > 0; i-- {
		res[i-1] += res[i] / 10
		res[i] = res[i] % 10
	}

	if res[0] == 0 {
		res = res[1:]
	}

	resArr := make([]byte, len(res))
	for i := 0; i < len(res); i++ {
		resArr[i] = '0' + byte(res[i])
	}

	return string(resArr)
}

func main() {
	num1, num2 := "123", "456"

	mulRes := multiply(num1, num2)

	fmt.Println(mulRes)
}
