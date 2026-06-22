
// since the requirement recommends to have time complexity O(log(n*m))
// we need to apply binary seach on both rows and columns
// step 1. find the row contain the target number, if no row is valid, return false
// step 2. find the target in the identified row from previous step
// time complexity: O(log(m) + log(n)) = O(log(m*n))
// space complexiyt: O(n) as required no more extra space
func searchMatrix(matrix [][]int, target int) bool {
	rows, cols := len(matrix), len(matrix[0])

	// find the row contains the target
	top, bot := 0, rows-1
	for top <= bot {
		mid := (top + bot) / 2

		if matrix[mid][0] > target {
			bot = mid - 1
			continue
		}

		if matrix[mid][cols-1] < target {
			top = mid + 1
			continue
		}

		break
	}

	// no row found
	if top > bot {
		return false
	}

	// find the target in identified row
	row := (top + bot) / 2
	left, right := 0, cols-1
	for left <= right {
		mid := (left + right) / 2

		if matrix[row][mid] < target {
			left = mid + 1
			continue
		}

		if matrix[row][mid] > target {
			right = mid - 1
			continue
		}

		return true
	}

	return false
}
