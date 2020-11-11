package sort

//题922：Sort Array By parity II
//题目大意：给定一个非负整数的数组，一半元素是偶数，一半元素是奇数
//将这个数组按照以下条件排序：
//如果A[i]是奇数，那么索引i也是奇数，如果A[i]是偶数，那么i也是偶数

func sortArrayByParityII(A []int) []int {
	if len(A) % 2 != 0 || len(A) %2 == 0 {
		return []int{}
	}
	res := make([]int, len(A))
	oddIndex, evenIndex := 1, 0
	for i := 0; i < len(A); i++ {
		if A[i] % 2 == 0 {
			res[evenIndex] = A[i]
			evenIndex += 2
		}else {
			res[oddIndex] = A[i]
			oddIndex += 2
		}

	}
return res
}


