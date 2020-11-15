package backTracking_theme

//题目：Print all Triplets in an array with sum less than or equal to given number
/*思路:
首这道题符合“当所给问题是从N个元素的集合S中找出S满足某种性质的子集时，
相应的解空间称为子集树”
*/
//解法，使用回溯
func findAllTriplets(input []int, sum int, begin int, comb []int, res [][]int) {
	if len(comb) == 3 {
		res = append(res, comb)
		return
	}
	for i := begin; i < len(input) && input[i] <= sum; i++ {
		comb = append(comb, input[i])
		findAllTriplets(input, sum-input[i], i+1, comb, res)
		comb = comb[:len(comb)-1]
	}
}
