package bit_manipulation_theme

//题268: Miss Number
//Given an array containing n distinct numbers taken from 0, 1, 2, ..., n, find the one that is missing from the array.
//Input: [3,0,1]
//Output: 2

//依次将数组中的元素异或，最后那个就是没有出现过的数
func findMissNumber(nums []int) int {
	xor, i := 0, 0
	for i = 0; i < len(nums); i++ {
		xor = xor ^ nums[i] ^ i
	}
	return xor ^ i
}
