// loop through the operations
// and push/pop element as the rules
// return the sum of elements of the stack at the end
func calPoints(operations []string) int {
	stack := make([]int, 0, len(operations))
	sum := 0
	for i := range operations {
		switch(operations[i]) {
			case "D":
				top := stack[len(stack)-1]
				stack = append(stack, 2*top)
				sum += 2 *top
			case "+":
				top1 := stack[len(stack)-1]
				top2 := stack[len(stack)-2]
				stack = append(stack, top1 + top2)
				sum += top1 + top2
			case "C":
				top := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				sum -= top
			default:
				num, err := strconv.Atoi(operations[i])
				if err != nil {
					panic(fmt.Sprintf("wrong operation: %s", operations[i]))
				}
				stack = append(stack, num)
				sum += num
		}
	}

	return sum
}
