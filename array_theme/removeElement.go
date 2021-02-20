package array_theme

// Given an array nums and a value val, remove all instances of that value in-place and return the new length.
// Do not allocate extra space for another array, you must do this by modifying the input array in-place with O(1) extra memory.
// The order of elements can be changed. It doesn't matter what you leave beyond the new length.

func removeElement(nums []int, val int) int {
	begin := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[begin] = nums[i]
			begin++
		}
	}
	return begin
	// nums = [1,2,1,3] val=1
	// step 1: nums[0]=1 == val=1  now i=0 nums=[1,2,1,3] begin=0
	// step 2: nums[1]=2 != val=1  now i=1 nums=[2,2,1,3] begin=1
	// step 3: nums[2]=1 == val=1  now i=2 nums=[2,2,1,3] begin=1
	// step 4: nums[3]=3 != val=1  now i=3 nums=[2,3,1,3] begin=2
}