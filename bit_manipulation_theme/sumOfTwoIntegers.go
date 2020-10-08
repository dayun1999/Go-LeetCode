package bit_manipulation_theme

//题371：Sum of Two Integers
//Calculate the sum of two integers a and b,
//but you are not allowed to use the operator + and -.
func getSum(a int, b int) int {
	if a == 0 {
		return b
	}
	if b == 0 {
		return a
	}
	//a&b << 1 是计算进位的
	//a^b 是计算不带进位的加法
	return getSum((a&b)<<1, a^b)
}
