func threeSum(nums []int) [][]int {
	sort.Slice(nums, func(i, j int) bool {return nums[i] < nums[j]})

	pos := make(map[[3]int]bool)
	ans := make([][]int, 0)
	for i := 0; i < len(nums); i++ {
		j := i + 1
		k := len(nums) - 1
		
		for j < k {
			if nums[j] + nums[k] + nums[i] == 0 {
				if _, ok := pos[[3]int{nums[i], nums[j], nums[k]}]; !ok {
					ans = append(ans,[]int{nums[i], nums[j], nums[k]})
					pos[[3]int{nums[i], nums[j], nums[k]}] = true
				}
				j++
				k--
			} else if nums[j] + nums[k] + nums[i] > 0 {
				k--
			} else {
				j++
			}
		}
	}
	return ans
}
