package palindrome

import (
	"fmt"
	"testing"
)

func TestPalindrome(t *testing.T) {
	s := "babad"

	dp := Palindrome(s)

	fmt.Println(dp)
}
