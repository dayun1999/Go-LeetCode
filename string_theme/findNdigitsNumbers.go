package string_theme

import "fmt"

//题目：Find all N-digit numbers with equal sum of digits at even and odd index
//题目大意：找到所有N位的数字，该数字满足奇数位的和 == 偶数位的和
//比如N = 3的时候，满足要求的所有数字为：
//110 121 132 143 ... 891 990

func findNdigitsNumbers(res []rune, N int, diff int) {
	if N > 0 {
		ch := '0'
		//特殊情况，数字不能以0开头
		if len(res) == 0 {
			ch = '1'
		}
		for ch <= '9' {
			absDifference := 0
			//更新奇数位与偶数位的diff
			if N&1 != 0 {
				//如果是奇数
				absDifference = diff + int(ch-'0')
			} else {
				//如果是偶数
				absDifference = diff - int(ch-'0')
			}
			findNdigitsNumbers(append(res, ch), N-1, absDifference)
			ch++
		}
	} else if N == 0 && abs(diff) == 0 {
		fmt.Println(string(res))
	}
}

func abs(a int) int {
	if a >= 0 {
		return a
	}
	return 0 - a
}
