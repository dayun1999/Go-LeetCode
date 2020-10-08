package array_theme

import "math"

//给定一个数组，要求在这个数组中找出 3 个数之和离 target 最近

//法一: 蛮力法
func threeSumClosest1(nums []int, target int) int {
	res, difference := 0, math.MaxInt16
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			for k := j + 1; k < len(nums); k++ {
				if difference > abs(nums[i]+nums[j]+nums[k]-target) {
					difference = abs(nums[i] + nums[j] + nums[k] - target)
					res = nums[i] + nums[j] + nums[k]
				}
			}
		}
	}
	return res
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
