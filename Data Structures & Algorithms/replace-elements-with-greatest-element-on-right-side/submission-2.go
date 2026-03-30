// approach 2: prefix sum
// maxVal=-1
// loop through the element from len(arr) - 1 to 0
// assign current element to temp variable
// set current element to maxVal
// update maxVal = max(maxVal, current element)
func replaceElements(arr []int) []int {
	m := -1
	for i := len(arr) - 1; i >= 0; i-- {
		c := arr[i]
		arr[i] = m
		m = max(m, c)
	}

	return arr
}
