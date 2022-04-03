package structures_test

import (
	"testing"

	"github.com/SuvorovSergey/go-leetcode/structures"
)

func TestNums2List(t *testing.T) {
	tests := []struct {
		name  string
		input []int
	}{
		{
			name:  "test1",
			input: []int{3, 5, 2, 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := structures.Nums2List(tt.input)
			r := structures.List2Nums(l)

			for i, v := range tt.input {
				if v != r[i] {
					t.Errorf("nums2list() = %v, want %v", r, tt.input)
				}
			}
		})
	}
}
