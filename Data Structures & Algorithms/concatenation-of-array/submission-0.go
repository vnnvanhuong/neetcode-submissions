// one naive approach is to create a new array with size 2*n
// loop through element of nums and fill the elements for new array
// newArray[i] = nums[i]; newArray[i+n] = nums[i]
// time complexity: O(N)
// space complexity: O(N)
func getConcatenation(nums []int) []int {
	return append(nums, nums...)
}
