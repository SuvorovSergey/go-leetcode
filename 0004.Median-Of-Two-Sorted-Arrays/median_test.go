package median_test

import (
	"testing"

	median "github.com/SuvorovSergey/go-leetcode/0004.Median-Of-Two-Sorted-Arrays"
)

func TestFindMedianSortedArrays(t *testing.T) {
	tests := []struct {
		name     string
		arg1     []int
		arg2     []int
		expected float64
	}{
		{
			name:     "test1",
			arg1:     []int{1, 3},
			arg2:     []int{2},
			expected: 2,
		},
		{
			name:     "test2",
			arg1:     []int{1, 2},
			arg2:     []int{3, 4},
			expected: 2.5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := median.FindMedianSortedArrays(tt.arg1, tt.arg2)
			if got != tt.expected {
				t.Errorf("FindMedianSortedArrays() = %v, want %v", got, tt.expected)
			}
		})
	}
}
