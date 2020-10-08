package bit_manipulation_theme

func reverseBits(nums []byte) {
	left, right := 0, len(nums)-1
	for i := 0; left < right; i++ {
		nums[left], nums[right] = nums[right], nums[left]
	}
}
