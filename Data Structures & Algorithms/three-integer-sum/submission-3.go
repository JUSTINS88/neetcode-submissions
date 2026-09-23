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
			if j - 1 != i && nums[j] == nums[j - 1] {
				j++
			} else if k != len(nums) - 1 && nums[k] == nums[k + 1] {
				k--
			} else if nums[j] + nums[k] + nums[i] == 0 {
				ans = append(ans,[]int{nums[i], nums[j], nums[k]})	
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
