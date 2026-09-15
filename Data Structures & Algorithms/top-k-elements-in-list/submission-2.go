func topKFrequent(nums []int, k int) []int {
	freq := make(map[int]int)
	bucket := make(map[int][]int)
	ans := make([]int, 0)
	for _, num := range nums {
		freq[num]++
	}

	for key, value := range freq {
		bucket[value] = append(bucket[value], key)
		
	}

	for i := len(nums); i >= 1; i-- {
		if _, ok := bucket[i]; ok {
			for j := 0; j < len(bucket[i]) && k > 0; j++ {
				ans = append(ans, bucket[i][j])
				k--
			}
		}
	}
	return ans
}
