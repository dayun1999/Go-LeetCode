package DynamicProgramming

//题53：Maximum Subarray
//Given an integer array nums, find the contiguous subarray (containing at least one number)
//which has the largest sum and return its sum.

//解题思路：设dp[i]为子数组[0,i]连续子数组的最大值之和
//则dp[i]=dp[i-1]+nums[i] (if dp[i-1]>0)  ||  dp[i]=dp[i-1] (if dp[i-1]<0)
func maxSubarray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	dp, res := make([]int, len(nums)), nums[0]
	dp[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		if dp[i-1] < 0 {
			dp[i] = nums[i]
		} else {
			dp[i] = dp[i-1] + nums[i]
		}
		res = max(res, dp[i])
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
