package array_theme

func maxArea(arr []int) int {
	max, left, right := 0, 0, len(arr)-1
	for left < right {
		height, width := 0, right-left
		if arr[left] < arr[right] {
			height = left
			left++
		} else {
			height = right
			right--
		}
		temp := height * width
		if temp > max {
			max = temp
		}
	}
	return max
}
