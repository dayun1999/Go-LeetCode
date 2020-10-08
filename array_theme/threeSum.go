package array_theme

func threeSum(target []int) [][]int {
	m := make(map[int]int)
	//find three elements whose sum is zero
	for i := 0; i < len(target); i++ {
		if _, ok := m[i]; ok {
			for j := i + 1; j < len(target); j++ {

			}
		}
		m[i] = target[i]
	}
	return nil
}

//[-2 -3 5 6 -1 -1 1 3]
//[-3 -2 -1 -1 1 3 5 6]
