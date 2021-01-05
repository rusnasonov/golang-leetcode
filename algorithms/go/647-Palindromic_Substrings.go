package main

func isPalindrome(s *string, start, end int) bool {
	for start < end {
		if (*s)[start] != (*s)[end-1] {
			return false
		}
		start++
		end--
	}
	return true
}

// Check All Substring. Simple solution
func countSubstringsDynamic(s string) int {
	n, ans := len(s), 0

	if (n <= 0) {
		return 0
	} 

	dp := make([][]bool, n)
	for i := range dp {
		dp[i] = make([]bool, n)
	}

	// Base case: single letter substrings
	for i := 0; i < n; i++ {
		dp[i][i] = true
		ans++
	} 
		

	// Base case: double letter substrings
	for i := 0; i < n - 1; i++ {
		dp[i][i + 1] = (s[i] == s[i + 1]);
		if dp[i][i+1] {
			ans++
		}
	}

	// All other cases: substrings of length 3 to n
	for len := 3; len <= n; len++ {
		i := 0
		for j := i + len - 1; j < n; i, j = i + 1, j + 1 {
			dp[i][j] = dp[i + 1][j - 1] && s[i] == s[j];
			if dp[i][j] {
				ans++
			}
		}
	}
	return ans;
}

