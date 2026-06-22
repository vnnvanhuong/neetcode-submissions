// since the recommende time complexity is O(nlogm)
// we need to apply binary search on m, where m is max number of bananas M in a pile
// it turns to find a x in range [1, M] where hours consuming with x is less or equal to h
// hour consuming is calculated as ceil(bananas / x)
func minEatingSpeed(piles []int, h int) int {
	lower, upper := 1, maxNumber(piles)
	answer := upper
	for lower <= upper {
		mid := lower + (upper - lower)/2
		if isSastified(piles, mid, h) {
			answer = mid
			upper = mid - 1
			continue
		}

		lower = mid + 1
	}

	return answer
}

func isSastified(piles []int, x, h int) bool {
	hours := 0
	for _, p := range piles {
		hours += (p + x - 1) / x
	}
	return hours <= h
}

func maxNumber(piles []int) int {
	m := 0
	for _, p := range piles {
		m = max(m, p)
	}
	return m
}