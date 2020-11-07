package DynamicProgramming

//题121: Best Time to buy and sell stock
//Input: [7,1,5,3,6,4]
//Output: 5
//Explanation: Buy on day 2 (price = 1) and sell on day 5 (price = 6), profit = 6-1 = 5.
//Not 7-1 = 6, as selling price needs to be larger than buying price.
func sellStock(nums []int) int {
	if len(nums) < 1 {
		return 0
	}
	min, maxProfit := nums[0], 0
	for i := 1; i < len(nums); i++ {
		if nums[i]-min > maxProfit {
			maxProfit = nums[i] - min
		}
		if nums[i] < min {
			min = nums[i]
		}
	}
	return maxProfit
}
