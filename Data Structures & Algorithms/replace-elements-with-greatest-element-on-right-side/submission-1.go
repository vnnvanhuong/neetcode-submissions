// approach 1: brute force. 
// for each element, find the max to its right, and replace it
func replaceElements(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		if i == len(arr) -1 {
			arr[i] = -1
			break
		}

		m := 0
		for j := i + 1; j < len(arr); j++ {
			m = max(m, arr[j])
		}
		arr[i]=m
	}

	return arr
}
