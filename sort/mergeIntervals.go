package sort

//题56：Merge Intervals
//题目大意：给定一个集合，合并所有重复的子数组

type Interval struct {
	Start int
	End   int
}

//合并的过程
func merge56(intervals []Interval) []Interval {
	if len(intervals) == 0 {
		return intervals
	}
	//排序
	quickSort(intervals, 0, len(intervals)-1)
	result := make([]Interval, 0)
	result = append(result, intervals[0])
	curIndex := 0
	for i := 1; i < len(intervals); i++ {
		//错误 if intervals[i].Start < intervals[curIndex].End {}
		if intervals[i].Start > result[curIndex].End {
			curIndex++
			//直接添加即可
			result = append(result, intervals[i])
		} else {
			//也就是intervals[i].Start < result.End,需要合并
			result[curIndex].End = max(intervals[i].End, result[curIndex].End)
		}

	}

	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//快速排序，将集合的每个数组里面的每个元素排序
func quickSort(a []Interval, low, high int) {
	//首先定义退出递归的条件
	if low >= high {
		return
	}
	//找到基准
	p := partitionSort(a, low, high)
	quickSort(a, low, p-1)
	quickSort(a, p+1, high)
}

//快速排序又叫分区交换函数，这是分区

//algorithm partition(A, lo, hi) is
//pivot := A[hi]
//i := lo
//for j := lo to hi do
//if A[j] < pivot then
//swap A[i] with A[j]
//i := i + 1
//swap A[i] with A[hi]
//return i

func partitionSort(a []Interval, low, high int) int {
	//首先定义一个基准
	pivot := a[high]
	i := low - 1
	for j := low; j < high; j++ {
		if a[j].Start < pivot.Start || (a[j].Start == pivot.Start &&
			a[j].End < pivot.End) {
			i++
			//交换
			a[i], a[j] = a[j], a[i]
		}
	}
	//a[i], a[high] = a[high], a[i]
	a[i+1], a[high] = a[high], a[i+1]
	return i
}

//自己写的
//func mergeIntervals(nums [][]int) [][]int {
//	temp, result := []int{}, [][]int{}
//	for i := 1; i < len(nums); i++ {
//		if nums[i-1][len(nums[i-1])-1] >= nums[i][0] {
//			temp = append(temp, nums[i-1][0])
//			temp = append(temp, nums[i][len(nums[i])-1])
//		} else {
//			temp = nums[i]
//		}
//
//		result = append(result, temp)
//	}
//
//	return result
//}
