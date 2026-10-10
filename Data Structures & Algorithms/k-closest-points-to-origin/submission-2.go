func kClosest(points [][]int, k int) [][]int {
	ans := make(map[int][][]int)
	var pos []int
	var fin [][]int

	for _, val := range points {
		sq := val[0] * val[0] + val[1] * val[1]
		if _, ok := ans[sq]; !ok {
			pos = append(pos, sq)
		}
		ans[sq] = append(ans[sq], val)
	}

	sort.Ints(pos)

	i := 0
	for k > 0 {
		for _, val := range ans[pos[i]] {
			fin = append(fin, val)
			k--
			if k == 0 {
				break
			}
		}
		i++
	}

	return fin
}
