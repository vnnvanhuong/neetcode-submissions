// Definition for a pair.
// type Pair struct {
//     Key   int
//     Value string
// }

type Solution struct {

}

func NewSolution() *Solution {
	return &Solution{}
}

func QuickSort(pairs []Pair) []Pair {
	return qs(pairs, 0, len(pairs)-1)
}

func qs(pairs []Pair, start, end int) []Pair{
	// base case
	if end - start <= 0 {
		return pairs
	}

	pivot := pairs[end]
	left := start

	// swap
	for i := start; i < end; i++ {
		if pairs[i].Key < pivot.Key {
			pairs[left], pairs[i] = pairs[i], pairs[left]
			left++
		}
	}

	// set pivot at left pointer
	pairs[end], pairs[left] = pairs[left], pivot

	qs(pairs, start, left-1)
	qs(pairs, left+1, end)

	return pairs
}
