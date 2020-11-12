package sort

//题75 Sort Colors
//题目大意：给定一个元素只有红色、白色和蓝色的颜色数组，将数组排序，红色 <-- 白色 <-- 蓝色
//其中，用0 1 2 分别表示 红色 白色 蓝色, 要求就地排序(不能新建数组)

//我的思路：
//遍历原始数组，将元素0 1 2 分别放进3个数组，最后再按顺序添加到一个大的数组中并且返回
//估计时间复杂度为O(n) 空间复杂度为O(K)
//我的解法--违反了就地排序：
func sortColors(colors []int) []int {
	res, temp0, temp1, temp2 := make([]int, 0, len(colors)), make([]int, 0), make([]int, 0), make([]int, 0)
	for _, v := range colors {
		switch v {
		case 0:
			temp0 = append(temp0, v)
			break
		case 1:
			temp1 = append(temp1, v)
			break
		case 2:
			temp2 = append(temp2, v)
			break
		}
	}
	res = append(res, temp0...)
	res = append(res, temp1...)
	res = append(res, temp2...)
	return res
}

//高要求：能否一次遍历，空间复杂度为O(1)呢？
//高手思路：由于数组只出现3种元素0 1 2，所以可以用游标移动来控制顺序
func sortColors2(colors []int) {
	if len(colors) == 0 {
		return
	}
	r, w, b := 0, 0, 0
	for _, num := range colors {
		if num == 0 {
			colors[b] = 2
			b++
			colors[w] = 1
			w++
			colors[r] = 0
			r++
		} else if num == 1 {
			colors[b] = 2
			b++
			colors[w] = 1
			w++
		} else if num == 2 {
			b++
		}
	}
}
