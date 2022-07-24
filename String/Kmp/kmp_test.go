package kmp

import (
	"fmt"
	"testing"
)

func TestKmp(t *testing.T) {
	p := "abcabe"
	s := "abcabfabcabef"

	index := IndexOf(s, p)

	fmt.Printf("find %s in %s, index is %d", p, s, index)

}
