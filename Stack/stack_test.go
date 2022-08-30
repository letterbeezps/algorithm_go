package stack

import (
	"fmt"
	"testing"
)

func TestStack(t *testing.T) {
	test1 := "()[]{}"
	fmt.Println(isValids(test1))
}

// go test -run TestValidPushPop
func TestValidPushPop(t *testing.T) {
	pushed := []byte("abcde")
	poped := []byte("decba")
	fmt.Println(validateStackSequences(pushed, poped))
}
