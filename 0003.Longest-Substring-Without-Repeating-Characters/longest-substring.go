package longestsubstring

// Runtime: 16 ms, faster than 41.58% of Go online submissions 
// Memory Usage: 3.1 MB, less than 57.39% of Go online submissions 
func LengthOfLongestSubstring(s string) int {
	m := make(map[byte]int)
	longest, left, right := 0, 0, 0

	for right < len(s) {
		if i, ok := m[s[right]]; ok && i >= left {
			left = i + 1
		}

		m[s[right]] = right
		right++

		if longest < (right - left) {
			longest = right - left
		}
	}

	return longest
}
