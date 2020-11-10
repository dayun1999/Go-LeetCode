package sort

//题242: Valid Anagram
//题目大意：给定两个string s 和 t, 判断是否是相同字符异序词

//优化解法
func isValidAnagram(s, t string) bool {
	//定义一个字母表
	alphabet := make([]int, 26)
	sBytes := []byte(s)
	tBytes := []byte(t)
	if len(sBytes) != len(tBytes) {
		return false
	}
	for i := 0; i < len(sBytes); i++ {
		alphabet[sBytes[i]-'a']++
	}
	for i := 0; i < len(tBytes); i++ {
		alphabet[tBytes[i]-'a']--
	}
	for i := 0; i < len(sBytes); i++ {
		if alphabet[i] != 0 {
			return false
		}
	}
	return true
}

//自己做的
//func isValidAnagram(s, t string) bool {
//	if len(s) != len(t) {
//		return false
//	}
//	res1 := sortString(s)
//	res2 := sortString(t)
//	for i := 0; i < len(res1); i++ {
//		if res1[i] != res2[i] {
//			return false
//		}
//	}
//
//	return true
//}
//
//func sortString(t string) string {
//	target := []byte(t)
//	for i := 0; i < len(target); i++ {
//		for j := i; j < len(target); j++ {
//			if target[i] > target[j] {
//				//交换
//				target[i], target[j] = target[j], target[i]
//			}
//		}
//	}
//	return string(target)
//}
