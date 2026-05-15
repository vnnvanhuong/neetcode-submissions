func kClosest(points [][]int, k int) [][]int {
	quickSort(points, 0, len(points)-1)
	res := make([][]int, 0, k)
	for i := 0; i < k; i++ {
		res = append(res, points[i])
	}

	return res
}

func quickSort(points [][]int, start, end int) [][]int {
	if end - start <= 0 {
		return points
	}

	pivot, left := points[end], start
	for i := start; i < end; i++ {
		if distance(points[i]) < distance(pivot) {
			points[left], points[i] = points[i], points[left]
			left++
		}
	}

	points[end], points[left] = points[left], pivot
	quickSort(points, start, left-1)
	quickSort(points, left+1, end)
	return points
}

func distance(p []int) float64 {
	return math.Sqrt(float64(p[0]*p[0] + p[1]*p[1]))
}

// quicksort
// calculate the distance by Euclideab formular 
// pick k points from sorted points