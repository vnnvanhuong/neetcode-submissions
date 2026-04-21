func climbStairs(n int) int {
	memo := make(map[int]int)
	var dp func(int) int
	dp = func(i int) int {
		if i > n {
			return 0
		}
		
		if i == n {
			return 1
		}

		if _, ok := memo[i]; ok {
			return memo[i]
		}

		t := dp(i+1) + dp(i+2)
		memo[i] = t

		return memo[i]
	}

	return dp(0)
}
