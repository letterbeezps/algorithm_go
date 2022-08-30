package stack

import (
	"fmt"
	"testing"
)

func TestStack(t *testing.T) {
	test1 := "()[]{}"
	fmt.Println(isValids(test1))
}
