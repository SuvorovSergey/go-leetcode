package addtwonumbers_test

import (
	"testing"

	addtwonumbers "github.com/SuvorovSergey/go-leetcode/0002.Add-Two-Numbers"
	"github.com/SuvorovSergey/go-leetcode/structures"
)

func TestAddTwoNumbers(t *testing.T) {
	tests := []struct {
		name     string
		input1   []int
		input2   []int
		expected []int
	}{
		{
			name:     "test1",
			input1:   []int{2, 4, 3},
			input2:   []int{5, 6, 4},
			expected: []int{7, 0, 8},
		},
		{
			name:     "test2",
			input1:   []int{0},
			input2:   []int{0},
			expected: []int{0},
		},
		{
			name:     "test3",
			input1:   []int{9, 9, 9, 9, 9, 9, 9},
			input2:   []int{9, 9, 9, 9},
			expected: []int{8, 9, 9, 9, 0, 0, 0, 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l1 := structures.Nums2List(tt.input1)
			l2 := structures.Nums2List(tt.input2)
			res := addtwonumbers.AddTwoNumbers(l1, l2)
			r := structures.List2Nums(res)

			if len(tt.expected) != len(r) {
				t.Errorf("AddTwoNumbers() = %v, want %v", r, tt.expected)
				return
			}

			for i, v := range tt.expected {
				if v != r[i] {
					t.Errorf("AddTwoNumbers() = %v, want %v", r, tt.expected)
				}
			}
		})
	}

}
