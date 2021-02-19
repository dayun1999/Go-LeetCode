package array_theme

// Given a binary array, find the maximum number of consecutive 1s in this array.

func findMaxConsecutiveOnes_Solution1(nums []int) int {
	start, end, res := -1, -1, 0
	//遍历数组
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			start = i
			end = start
			continue
		}
		end++
		res = max(res, end-start)
	}

	return res
}

func findMaxConsecutiveOnes_Solution2(nums []int) int {
	maxHere, res := 0, 0
	for _, n := range nums {
		if n == 0 {
			maxHere = 0
		} else {
			maxHere++
		}
		res = max(res,maxHere )
		
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
