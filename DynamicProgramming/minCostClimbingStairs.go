package DynamicProgramming

//题746：Min Cost Climbing Stairs
//题目大意：一次爬1或者2层楼梯，每一层台阶对应一个非负整数的cost，算出到达顶点
//的最小值
//动态规划状态转移方程为：dp[i] = cost[i] + min{dp[i-1], dp[i-2]}
func minCostClimbingStairs(cost []int) int {
	if len(cost) < 2 || len(cost) > 1000 {
		return 0
	}
	dp := make([]int, len(cost))
	dp[0], dp[1] = cost[0], cost[1]
	for i := 2; i < len(cost); i++ {
		dp[i] = cost[i] + min(dp[i-2], dp[i-1])
	}
	return min(dp[len(cost)-1], dp[len(cost)-2])
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
