package median

// https://medium.com/@hazemu/finding-the-median-of-2-sorted-arrays-in-logarithmic-time-1d3f2ecbeb46
// https://www.geeksforgeeks.org/median-of-two-sorted-arrays-of-different-sizes/
// https://olympiads.ru/moscow/2008/team/problems/medians_razbor.pdf

// Runtime: 8 ms, faster than 96.46% of Go online submissions for Median of Two Sorted Arrays.
// Memory Usage: 5 MB, less than 91.57% of Go online submissions for Median of Two Sorted Arrays.
func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {

	if len(nums1) > len(nums2) {
		return FindMedianSortedArrays(nums2, nums1)
	}

	k := (len(nums1) + len(nums2) + 1) >> 1
	nums1MinCount := 0
	nums1MaxCount := len(nums1)
	nums1Count := 0
	nums2Count := 0

	for nums1MinCount <= nums1MaxCount {
		nums1Count = nums1MinCount + (nums1MaxCount-nums1MinCount)>>1
		nums2Count = k - nums1Count

		if nums1Count > 0 && nums1[nums1Count-1] > nums2[nums2Count] {
			nums1MaxCount = nums1Count - 1
		} else if nums1Count < len(nums1) && nums2[nums2Count-1] > nums1[nums1Count] {
			nums1MinCount = nums1Count + 1
		} else {

			break

		}

	}

	ml, mr := 0, 0

	if nums1Count == 0 {
		ml = nums2[nums2Count-1]
	} else if nums2Count == 0 {
		ml = nums1[nums1Count-1]
	} else {
		if nums2[nums2Count-1] > nums1[nums1Count-1] {
			ml = nums2[nums2Count-1]
		} else {
			ml = nums1[nums1Count-1]
		}
	}

	if (len(nums1)+len(nums2))%2 != 0 {
		return float64(ml)
	}

	if nums1Count == len(nums1) {
		mr = nums2[nums2Count]
	} else if nums2Count == len(nums2) {
		mr = nums1[nums1Count]
	} else {
		if nums2[nums2Count] < nums1[nums1Count] {
			mr = nums2[nums2Count]
		} else {
			mr = nums1[nums1Count]
		}
	}

	return float64((ml + mr)) / 2
}
