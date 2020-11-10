package DynamicProgramming

//题198：House Robber
//题目大意：作为一名抢劫犯，如何在一条街上不抢相邻的店铺但是又能获得最大额度的金钱呢

func houseRobber(shops []int) int {
	length := len(shops)
	if length == 0 {
		return 0
	}
	if length == 1 {
		return shops[0]
	}
	moneyFromOdd, moneyFromEven := 0, 0

	for i := 0; i < length; i++ {
		if i%2 == 0 {
			moneyFromEven += shops[i]
			moneyFromEven = max(moneyFromEven, moneyFromOdd)
		} else {
			moneyFromOdd += shops[i]
			moneyFromOdd = max(moneyFromOdd, moneyFromEven)
		}
	}
	return max(moneyFromOdd, moneyFromEven)
}
