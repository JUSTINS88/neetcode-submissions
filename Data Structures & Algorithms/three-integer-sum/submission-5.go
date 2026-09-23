func threeSum(nums []int) [][]int {
	sort.Slice(nums, func(i, j int) bool {return nums[i] < nums[j]})

	ans := make([][]int, 0)

	for i := 0; i < len(nums); i++ {
		j := i + 1
		k := len(nums) - 1
		if i > 0 && nums[i] == nums[i - 1] {
			continue
		}

		for j < k {
			if nums[j] + nums[k] + nums[i] == 0 {
				ans = append(ans,[]int{nums[i], nums[j], nums[k]})	
				j++
				k--
				for nums[j] == nums[j - 1] && j < k {j++}
				for nums[k] == nums[k + 1] && j < k {k--}
			} else if nums[j] + nums[k] + nums[i] > 0 {
				k--
			} else {
				j++
			}
		}
	}
	return ans
}
