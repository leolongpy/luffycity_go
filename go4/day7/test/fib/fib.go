package fib

func Fib(n int) int {
	if n < 2 {
		return n
	}
	return Fib(n-1) + Fib(n-2)
}

func Fib2(n int) int {
	dp := make([]int, n+1)
	dp[0] = 0
	dp[1] = 1
	for i := 2; i < n; i++ {
		dp[i] = dp[i-2] + dp[i-1]
	}
	return dp[n]
}
