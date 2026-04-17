// the order of students does not matter
// but the order of sandwiches matters
// count preferences of students
// loop through sandwiches
// reduce the preferences
// return the remaining students
// students = [1,1,1,0,0,1], sandwiches = [1,0,0,0,1,1]
// {1: 4, 0: 2}
// {1:3, 0: 2}
// {1:3, 0: 1}
// {1:3, 0: 0}
func countStudents(students []int, sandwiches []int) int {
    res := len(students)
	count := make(map[int]int, 2)
	for _, s := range students {
		count[s]++
	}

	for _, s := range sandwiches {
		if count[s] == 0 {
			return res
		}

		res--
		count[s]--
	}

	return res
}