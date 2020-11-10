package DynamicProgramming

//题942：DI String Match
//题目大意：给定一个只含有字符"I"或者"D"的字符串S，设定N=len(S)
//建立一个数组A=[0,1,2,3......N],满足以下条件
//if S[i]==“I“  A[i] < A[i+1]
//if S[i]=="D"  A[i] > A[i+1]

func diStringMatch(s string) []int {
	result, maxNum, minNum, index := make([]int, len(s)+1), len(s), 0, 0
	for _, ch := range s {
		if ch == 'I' {
			result[index] = minNum
			minNum++
		} else {
			result[index] = maxNum
			maxNum--
		}
		index++
	}
	result[index] = minNum
	return result
}
