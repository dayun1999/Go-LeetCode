package array_theme

// Given two sorted integer arrays nums1 and nums2, merge nums2 into nums1 as one sorted array.
// The number of elements initialized in nums1 and nums2 are m and n respectively.
// You may assume that nums1 has a size equal to m + n such that it has enough space to hold additional elements from nums2.
// nums1.length == n + m  nums2.length = n

func merge(nums1 []int, m int, nums2 []int, n int) {
	// 利用双指针
	i, j, k := m-1, n-1, n+m-1
	for i >= 0 && j >= 0 {
		if nums1[i] < nums2[j] {
			nums1[k] = nums2[j]
			k--
			j--
		} else {
			nums1[k] = nums1[i]
			k--
			i--
		}
	}

	// nums1 = [4 5 6 0 0 0]
	// nums2 = [1 2 3]
	// step1: nums1 = [4 5 6 0 0 6] 此时i=1 k=4 j=2
	// step2: nums1 = [4 5 6 0 5 6] 此时i=0 k=3 j=2
	// step3: nums1 = [4 5 6 4 5 6] 此时i=-1 k=2 j=2 不满足for循环的条件, 跳出循环
	// 因此我们需要再次判断
	for j >= 0 {
		nums1[k] = nums2[j]
		k--
		j--
	}
}
