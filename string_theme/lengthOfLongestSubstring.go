package string_theme

//Input: "abcabcbb"
//Output: 3
//Explanation: The answer is "abc", with the length of 3.

//法1：滑动窗口
//pwwkew
func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	var freq [256]int //只针对完整版的ASCII码表(256个)
	result, left, right := 0, 0, -1

	for left < len(s) {
		//右窗口一直像右滑动，如果遇见相同的字符那么左边的窗口就往右移动
		if right+1 < len(s) && freq[s[right+1]-'a'] == 0 {
			//将新出现的字母所在的位置+1
			freq[s[right+1]-'a']++
			right++
		} else {
			//如果遇见重复的字符，就-1，知道为0为止
			freq[s[left]-'a']--
			left++
		}
		result = max(result, right-left+1)
	}
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//自己写的
//func lengthOfLongestSubstring(str string) int {
//	var left, right , result int
//	length := len(str)
//	if length < 1 {
//		return 0
//	}
//	for i:=0; i<length && right < length;i++ {
//		right++
//		if str[i] != str[right] {
//			right++
//			result++
//		}else{
//			left = right
//		}
//	}
//	return result
//}
