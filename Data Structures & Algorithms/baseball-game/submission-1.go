// loop through the operations
// and push/pop element as the rules
// return the sum of elements of the stack at the end
func calPoints(operations []string) int {
	stack := make([]int, 0, len(operations))
	sum := 0
	val := 0
	for i := range operations {
		switch(operations[i]) {
			case "D":
				val = 2*stack[len(stack)-1]
				stack = append(stack, val)
				sum += val
			case "+":
				val = stack[len(stack)-1] + stack[len(stack)-2]
				stack = append(stack, val)
				sum += val
			case "C":
				val = stack[len(stack)-1]
				sum -= val
				stack = stack[:len(stack)-1]
			default:
				val, err := strconv.Atoi(operations[i])
				if err != nil {
					panic(fmt.Sprintf("wrong operation: %s", operations[i]))
				}
				stack = append(stack, val)
				sum += val
		}
	}

	return sum
}
