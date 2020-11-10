package DynamicProgramming

//题1025：Divisor Game
//题目大意：有一个数字N，两位玩家Alice 和 Bob需要选一个x,满足以下条件：
//	0  < x < N 并且 N % x == 0
//将N - x 替换 N,直到由一方找不到这个x就算输,Alice先玩，如果Alice赢了就返回True

func divisorGame(n int) bool {
	return n%2 == 0
}
