package array_theme

// Given an array nums, write a function to move all 0's to the end of it while maintaining the relative order of the non-zero elements.
// 要求是“就地(in-place)”解决, 而且非零元素的相对位置保持不变

// 方法: 滚雪球模式
func moveZeros(nums []int) {
	snowBallSize := 0
	for i := 0; i < len(nums); i++ {
		// 遇见一个0,就将"雪球“滚打
		if nums[i] == 0 {
			snowBallSize++
		} else if snowBallSize > 0 {
			nums[i], nums[i-snowBallSize] = nums[i-snowBallSize], nums[i]
		}
	}
}