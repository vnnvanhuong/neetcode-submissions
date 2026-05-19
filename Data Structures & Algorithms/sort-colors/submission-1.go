// since the elements of array are knew, we can use bucket sort
// need a pointer to keep track the current position to insert
// need a counter to count the occurences of each 0, 1, 2
func sortColors(nums []int) {
	low, mid, high := 0, 0, len(nums)-1
	for mid <= high {
		if nums[mid] == 0 {
			nums[low], nums[mid] = nums[mid], nums[low]
			low++
			mid++
			continue
		}

		if nums[mid] == 1 {
			mid++
			continue
		}

		nums[mid], nums[high] = nums[high], nums[mid]
		high--
	}
}
