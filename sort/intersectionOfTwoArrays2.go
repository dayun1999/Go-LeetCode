package sort

//题350：Intersection of two arrays II
//题目大意：找出两个数组的交集(不去重)

func findIntersectionII(nums1, nums2 []int) []int {
	/*思路
	和上一题类似，只不过这个不去重，所以每次比较不需要删除map中的元素
	*/
	m, result := make(map[int]int), []int{}
	for i := 0; i < len(nums1); i++ {
		m[nums1[i]]++
	}
	for i := 0; i < len(nums2); i++ {
		if _, ok := m[nums2[i]]; ok {
			result = append(result, nums2[i])
		}
	}
	return result
}
