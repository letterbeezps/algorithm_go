package main

import "fmt"

func main() {

	p := "abcabe"
	s := "abcabfabcabef"

	index := indexOf(s, p)

	fmt.Printf("find %s in %s, index is %d", p, s, index)
}

func indexOf(s, p string) int {
	next := make([]int, len(p))
	next[0] = -1

	j := -1
	for i := 1; i < len(p); i++ {
		for j != -1 && p[i] != p[j+1] {
			j = next[j]
		}
		if p[i] == p[j+1] {
			j++
		}
		next[i] = j
	}

	fmt.Println("p的next数组：")
	fmt.Println(next)

	j = -1
	for i := 0; i < len(s); i++ {
		for j != -1 && s[i] != p[j+1] {
			j = next[j]
		}
		if s[i] == p[j+1] {
			j++
		}

		if j == len(p)-1 {
			return i - j
		}
	}
	return -1
}
