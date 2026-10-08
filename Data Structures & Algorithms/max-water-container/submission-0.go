func maxArea(heights []int) int {
    maxx := 0
    l := 0
    r := len(heights) - 1

    for l < r {
        maxx = max(min(heights[l], heights[r]) * (r - l), maxx)
        if heights[l] <= heights[r] {
            l++
        } else {
            r--
        }
    }
    return maxx
}
