func hasDuplicate(nums []int) bool {
    appear := make(map[int]bool)
	for _, num := range nums {
		if _, ok := appear[num]; ok {
            return true
        }
        appear[num] = true
	}
	return false
}
