func findMaxConsecutiveOnes(nums []int) int {
	c, m := 0, 0
    for i := range nums {
        if nums[i]==0 {
            m=max(m, c)
            c = 0
            continue
        }

        c++
    }

    return max(m, c)
}
