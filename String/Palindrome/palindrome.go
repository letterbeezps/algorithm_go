package palindrome

// 计算字符串s内部的全部回文子串
// 得到一个数组 dp [][]bool
// dp[i][j] 表示字符串 s[i:j+1]是否为回文串
// 依此为拓展可以解决很多与回文串有关的问题
func Palindrome(s string) [][]bool {
	dp := make([][]bool, len(s))
	for i := 0; i < len(s); i++ {
		dp[i] = make([]bool, len(s))
		dp[i][i] = true
	}

	for i := len(s) - 1; i >= 0; i-- {
		for j := i + 1; j < len(s); j++ {
			dp[i][j] = (s[i] == s[j]) && ((j-i < 3) || dp[i+1][j-1])
		}
	}

	return dp
}
