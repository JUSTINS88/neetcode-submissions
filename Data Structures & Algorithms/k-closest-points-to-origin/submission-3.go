func kClosest(points [][]int, k int) [][]int {
	var ans [][]int
	sort.Slice(points, func (i, j int) bool {
		x := points[i][0] * points[i][0] + points[i][1] * points[i][1] 
		y := points[j][0] * points[j][0] + points[j][1] * points[j][1]
		return x < y
	})

	for i := 0; i < len(points) && k > 0; i++ {
		ans = append(ans, points[i])
		k--
	}
	return ans
}
