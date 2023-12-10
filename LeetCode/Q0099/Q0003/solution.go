package Q0003

func lengthOfLongestSubstring(s string) int {
	m := make(map[byte]int)
	ans := 0
	start := 0
	for end := 0; end < len(s); end++ {
		right := s[end]
		m[right] = m[right] + 1
		for m[right] > 1 {
			left := s[start]
			m[left] = m[left] - 1
			start++
		}
		ans = max(ans, end-start+1)
	}
	return ans
}
