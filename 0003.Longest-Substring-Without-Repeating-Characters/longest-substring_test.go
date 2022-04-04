package longestsubstring_test

import (
	"testing"

	longestsubstring "github.com/SuvorovSergey/go-leetcode/0003.Longest-Substring-Without-Repeating-Characters"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int
	}{
		{
			name:     "test1",
			input:    "abcabcbb",
			expected: 3,
		},
		{
			name:     "test2",
			input:    "bbbb",
			expected: 1,
		},
		{
			name:     "test3",
			input:    "pwwkew",
			expected: 3,
		},
		{
			name:     "test4",
			input:    "abba",
			expected: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := longestsubstring.LengthOfLongestSubstring(tt.input)

			if tt.expected != r {
				t.Errorf("LengthOfLongestSubstring() = %v, want %v", r, tt.expected)
			}
		})
	}
}
