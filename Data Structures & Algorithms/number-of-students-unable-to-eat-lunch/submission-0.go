// how about follw exactly the same description, but how to know when to stop the loop
// 1. empty
// 2. not empty, using count
func countStudents(students []int, sandwiches []int) int {
	skipped := 0
	for len(students) > 0 {
		if skipped == len(students) {
			return len(students)
		}

		if students[0] == sandwiches[0] {
			students = students[1:]
			sandwiches = sandwiches[1:]
			skipped = 0
			continue
		}

		students = append(students[1:], students[0])
		skipped++
	}

    return 0
}