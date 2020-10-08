package array_theme

func twoSum(arr []int, target int) []int {
	//set a map whose key = value & value = index
	m := make(map[int]int)
	for i := 0; i < len(arr); i++ {
		anotherPart := target - arr[i]
		if _, ok := m[anotherPart]; ok {
			return []int{m[anotherPart], i}
		}
		m[anotherPart] = i
	}
	return nil
}
