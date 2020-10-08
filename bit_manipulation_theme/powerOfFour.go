package bit_manipulation_theme

//题342:Power of Four
//Given an integer (signed 32 bits), write a function to check whether it is a power of 4.
//Input: 16
//Output: true

//法1：数论，证明(4^n - 1) % 3 == 0
func isPowerOfFour(num int) bool {
	return num > 0 && (num-1)%3 == 0
}

//法2：循环
func isPowerOfFour01(num int) bool {
	for num%4 == 0 && num >= 4 {
		num = num / 4
		if num == 1 {
			return true
		}
	}
	return num == 1
}
