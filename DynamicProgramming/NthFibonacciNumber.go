package DynamicProgramming

//题目：Program to find n’th Fibonacci number

func findNthFibNumber(n int) int {
	/*思路
	由bottom up可以直接求得*/
	if n <= 1 {
		return n
	}
	lookup := make([]int, n)
	lookup[1] = 1
	lookup[0] = 1
	for i := 2; i < n; i++ {
		lookup[i] = lookup[i-1] + lookup[i-2]
	}
	return lookup[n-1]
}
