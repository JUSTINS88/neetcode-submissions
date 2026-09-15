func hasDuplicate(nums []int) bool {
    appear := make(map[int]int)
	duplicate := false
	for _, num := range nums {
		appear[num]++
		if appear[num] > 1 {
			duplicate = true
		}
	}
	return duplicate
}
