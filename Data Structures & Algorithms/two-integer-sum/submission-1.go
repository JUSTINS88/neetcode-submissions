func twoSum(nums []int, target int) []int {
    muncul := make(map[int]int)
    ans := make([]int, 2)
    for i := range nums {
        targeted := target - nums[i]
        if _, ok := muncul[targeted]; ok {
            ans[0] = muncul[targeted]
            ans[1] = i
            return ans
        } else {
            muncul[nums[i]] = i
        }
    }
    return ans
}