package solution2

func climbStairs(n int) int {
	return fib(n, 1, 1)
}

func fib(n, a, b int) int {
	if n == 0 {
		return a
	}
	return fib(n-1, b, a+b)
}
