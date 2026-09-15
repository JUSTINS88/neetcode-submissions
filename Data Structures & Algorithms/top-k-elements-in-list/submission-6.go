func topKFrequent(nums []int, k int) []int {
	freq := make(map[int]int)
	bucket := make([][]int, len(nums) + 1)
	ans := make([]int, k)
	for _, num := range nums {
		freq[num]++
	}

	for key, value := range freq {
		bucket[value] = append(bucket[value], key)
		
	}

	for i := len(nums); i >= 1; i-- {
		if len(bucket[i]) > 0 {
			for j := 0; j < len(bucket[i]) && k > 0; j++ {
				ans[k - 1] = bucket[i][j]
				k--
			}
		}
	}
	return ans
}
