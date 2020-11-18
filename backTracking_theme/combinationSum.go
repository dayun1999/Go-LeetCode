package backTracking_theme

import "sort"

//题目39: Combination Sum
//题目大意：给定一个正整数数组candidates，每个元素都不重复，给定一个target
//找出所有满足以下条件的唯一的数组
//candidates的元素的和 == target
//单个candidates的元素可以重复多次
//Input: candidates = [2,3,6,7], target = 7,
//A solution set is:
//[
//[7],
//[2,2,3]
//]

func CombinationSum(candidates []int, target int) [][]int {
	if len(candidates) == 0 {
		return [][]int{}
	}
	c, res := []int{}, [][]int{}
	//排序，从小到大
	sort.Ints(candidates)
	findCombinationSum(candidates, target, 0, c, &res)
	return res
}
func findCombinationSum(nums []int, target int, index int, c []int, res *[][]int) {
	if target <= 0 {
		if target == 0 {
			//找到符合条件的了
			//复制当前的c，并且加入res
			b := make([]int, len(c))
			copy(b, c)
			*res = append(*res, b)
		}
		return
	}
	for i := index; i < len(nums); i++ {
		//剪枝优化
		if nums[i] > target {
			break
		}
		c = append(c, nums[i])
		findCombinationSum(nums, target-nums[i], i, c, res)
		c = c[:len(c)-1] //减掉末尾
	}
}

//自己写的，错的
//func combinationSum(candidates []int, target int,temp []int, res [][]int) {
//	//递归结束的条件
//	if sum(temp) == target {
//		res = append(res, temp)
//		return
//	}
//	for i := 0; i < len(candidates); i++ {
//		tempSum := candidates[i] + sum(temp)
//		if tempSum < target {
//			combinationSum(candidates, target - tempSum, temp, res )
//		} else if tempSum == target {
//			temp = append(temp, candidates[i])
//			res = append(res, temp)
//			return
//		}
//	}
//}
//
//func sum(num []int) int {
//	res := 0
//	for i := range num {
//		res += num[i]
//	}
//	return res
//}
