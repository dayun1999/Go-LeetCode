package bit_manipulation_theme

//题目-136
//给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。
//找出那个只出现了一次的元素。要求算法时间复杂度是线性的，并且不使用额外的辅助空间。
//Input: [2,2,1]
//Output: 1

//利用异或的性质-相同为0，不同为1
//x^x=0
func singleNumber(nums []int) int {
	result := 0
	for i := 0; i < len(nums); i++ {
		result ^= nums[i] //0 ^ X = X
	}
	return result
}

//自己做的
//func findSingleNumber(nums []int) int {
//	res := 0
//	m := make(map[int]int) //key-element & value-times
//	for i := range nums {
//		m[nums[i]]++
//	}
//	for k, v := range m {
//		if v == 1 {
//			res = k
//		}
//	}
//	return res
//}
