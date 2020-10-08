package bit_manipulation_theme

//题231: Power of Two
//Given an integer, write a function to determine if it is a power of two.
//Input: 1
//Output: true
//Explanation: 2^0 = 1

//二进制位操作法
func isPowerOfTwo(num int) bool {
	return num > 0 && (num&(num-1) == 0)
}

//自己写的-错误
//func isPowerOfTwo(num int) bool {
//	if num <= 0 {
//		return false
//	}
//	if num == 1 {
//		return true
//	}
//	res := 0
//	for res != 0 {
//		if res & 1 != 1{  //一直是偶数
//			res = num % 2
//		}else{
//			return false
//		}
//	}
//	return true
//}
