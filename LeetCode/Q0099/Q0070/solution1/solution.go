package solution1

func climbStairs(n int) int {
	a, b := 1, 1
	for ; n > 0; n-- {
		a, b = b, a+b
	}
	return a
}
