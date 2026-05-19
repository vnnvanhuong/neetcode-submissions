// since the elements of array are knew, we can use bucket sort
// need a pointer to keep track the current position to insert
// need a counter to count the occurences of each 0, 1, 2
func sortColors(nums []int) {
	count := make([]int, 3)
	for _, num := range nums {
		count[num]++
	}

	p := 0
	for i, c := range count {
		for j := 0; j < c; j++ {
			nums[p] = i
			p++
		}
	}
}
