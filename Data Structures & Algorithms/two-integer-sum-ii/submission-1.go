func twoSum(numbers []int, target int) []int {
	i := 0
	j := len(numbers) - 1
	ans := make([]int, 2, 2)
	for i < j {
		if numbers[i] + numbers[j] == target {
			ans[0] = i + 1
			ans[1] = j + 1
			break
		} else if numbers[i] + numbers[j] > target {
			j--
		} else {
			i++
		}
	}
	return ans
}
