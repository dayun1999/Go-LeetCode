package DynamicProgramming

//题70：Climbing Stairs --经典的爬楼梯问题,每次只能爬1 或者 2 步
//You are climbing a stair case. It takes n steps to reach to the top.
//Each time you can either climb 1 or 2 steps. In how many distinct ways can you climb to the top?
//Note: Given n will be a positive integer.

//分析：
//cost[i]表示要爬到第i层所有的可能数
//cost[i] = cost[i-1] + cost[i-2] 即斐波那契数列
func climbingStairs(n int) int {
	if n <= 0 {
		return 0
	}
	dp := make([]int, n+1)
	dp[0], dp[1] = 1, 1
	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}
