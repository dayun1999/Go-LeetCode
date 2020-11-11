package sort

//题976：Largest Perimeter Triangle
//题目大意，给定一个非负整数数组，选定3个元素，组成满足周长最长的三角形
//[1,1,2,3,4,5,9]
func largestPerimeter(nums []int) int {
	if len(nums) < 3 {
		return -1
	}
	sortArray(nums)
	res := 0
	index := len(nums) - 1
	for index >= 2 {
		if nums[index] < nums[index-1]+nums[index-2] && nums[index-1] < nums[index]+nums[index-2] && nums[index-2] < nums[index-1]+nums[index] {
			res = nums[index] + nums[index-1] + nums[index-2]
		}
		index--
	}
	return res
}

func sortArray(target []int) {
	for i := 0; i < len(target); i++ {
		for j := i; j < len(target); j++ {
			if target[i] > target[j] {
				//交换
				target[i], target[j] = target[j], target[i]
			}
		}
	}
}
