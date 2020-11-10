package sort

//题349：Intersection of two arrays
//题目大意：找出两个数组的交集

func findIntersection(nums1, nums2 []int) []int {
	/*自己的思路：
	循环判断两个数组，其中一个数组的元素是否在另一个里面，有就append进结果数组
	时间复杂度大概O(n^2)
	*/
	/*优化的思路：
	将一个数组的所有元素放进一个字典，在试图添加另一个数组的所有元素进入字典
	if !ok ; then append into the result
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
