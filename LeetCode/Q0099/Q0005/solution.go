package Q0005

import "strings"

func longestPalindrome(s string) string {
	t := "^#" + strings.Join(strings.Split(s, ""), "#") + "#$"
	n := len(t)
	p := make([]int, n)
	c, r := 0, 0
	le := 1
	start := 0
	for i := 1; i < n-1; i++ {
		iM := 2*c - i
		p[i] = 0
		if i < r {
			p[i] = min(r-i, p[iM])
		}
		for t[i+1+p[i]] == t[i-1-p[i]] {
			p[i]++
		}
		if i+p[i] > r {
			c = i
			r = i + p[i]
		}
		if p[i] > le {
			le = p[i]
			start = (i - p[i]) / 2
		}
	}
	return s[start : start+le]
}
