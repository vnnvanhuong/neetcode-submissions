/** 
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is higher than the picked number
 *			      1 if num is lower than the picked number
 *               otherwise return 0
 * func guess(num int) int;
 */

func guessNumber(n int) int {
    low, high := 1, n
	for low <= high {
		pick := low + (high-low)/2
		result := guess(pick)
		if result > 0 {
			low = pick + 1
			continue
		}

		if result < 0 {
			high = pick - 1
			continue
		}
		return pick
	}

	return -1 // never happen
}
