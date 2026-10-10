func kClosest(points [][]int, k int) [][]int {
	sort.Slice(points, func (i, j int) bool {
		x := points[i][0] * points[i][0] + points[i][1] * points[i][1] 
		y := points[j][0] * points[j][0] + points[j][1] * points[j][1]
		return x < y
	})

	return points[:k]
}
