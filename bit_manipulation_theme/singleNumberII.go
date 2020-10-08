package bit_manipulation_theme

//137题:
//给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现了三次。
//找出那个只出现了一次的元素。要求算法时间复杂度是线性的，并且不使用额外的辅助空间。
//Input: [2,2,3,2]
//Output: 3

func singleNumberII(nums []int) int {
	ones, twos := 0, 0
	for i := 0; i < len(nums); i++ {
		ones = (ones ^ nums[i]) & ^twos
		twos = (twos ^ nums[i]) & ^ones
	}
	return ones
}
