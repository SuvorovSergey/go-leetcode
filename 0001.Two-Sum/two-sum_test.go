package twosum_test

import (
	"reflect"
	"testing"

	twosum "github.com/SuvorovSergey/go-leetcode/0001.Two-Sum"
)

func TestTwoSum(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}

	type test struct {
		name     string
		args     args
		expected []int
	}

	tests := []test{
		{
			name: "test1",
			args: args{
				nums:   []int{2, 7, 11, 15},
				target: 9,
			},
			expected: []int{0, 1},
		},
		{
			name: "test2",
			args: args{
				nums:   []int{3, 2, 4},
				target: 6,
			},
			expected: []int{1, 2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := twosum.TwoSum(tt.args.nums, tt.args.target)
			if !reflect.DeepEqual(tt.expected, res) {
				t.Errorf("twoSum() = %v, expected %v", res, tt.expected)
			}
		})
	}
}
