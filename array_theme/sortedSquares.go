package array_theme

// Given an integer array nums sorted in non-decreasing order,
// return an array of the squares of each number sorted in non-decreasing order.
func sortedSquares(nums []int) []int {
	// 利用双指针
	i, j := 0, len(nums)-1
	res := make([]int, len(nums))
	for p := len(nums)-1; p >= 0; p-- {
		if abs(nums[i]) < abs(nums[j]) {
			res[p] = square(nums[j])
			j--
		} else {
			res[p] = square(nums[i])
			i++
		}
	}
	return res
}

func square(a int) int {
	return a * a
}