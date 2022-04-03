package leetcode

// Complecity O(n)
// Runtime: 14 ms, faster than 42.80% of Go online submissions for Two Sum.
// Memory Usage: 4.3 MB, less than 53.37% of Go online submissions for Two Sum.
func TwoSum(nums []int, target int) []int {
	buf := make(map[int]int)
	for i, v := range nums {
		if idx, ok := buf[target-v]; ok {
			return []int{idx, i}
		}
		buf[v] = i
	}
	return []int{}
}

// Complecity O(n^2)
func TwoSumVersion1(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}

	return []int{}
}
