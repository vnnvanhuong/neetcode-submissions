func combinationSum(nums []int, target int) [][]int {
    result := [][]int{}
	total := 0
	cur := []int{}

	var dfs func(i int)
	dfs = func(i int) {
		if total == target {
			temp := make([]int, len(cur))
			copy(temp, cur)
			result = append(result, temp)
			return
		}
		
		if i >= len(nums) || total > target {
			return
		}

		cur = append(cur, nums[i])
		total += nums[i]
		dfs(i) // reuse nums[i]

		cur = cur[:len(cur)-1]
		total -= nums[i]
		dfs(i+1) // skip nums[i]
	}

	dfs(0)

	return result
}
