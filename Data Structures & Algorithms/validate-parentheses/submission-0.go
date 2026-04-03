// loop through the elements
// if open bracket, put it into stack
// if close bracket, compare with popped element from stack if same type
// return true if size stack is empty
// time: O(N), space: O(N)
func isValid(s string) bool {
	stack := make([]byte, 0, len(s))

	for i := range s {
		if isOpen(s[i]) {
			stack = append(stack, s[i])
			continue
		}

		if len(stack) == 0 {
			// first character is a closing bracket
			return false
		}

		top := stack[len(stack)-1]
		if !isSameType(s[i], top) {
			return false
		}

		stack = stack[:len(stack)-1]
	}


    return len(stack) == 0
}

func isOpen(c byte) bool {
	return c == '(' || c == '{' || c == '['
}

func isSameType(s, d byte) bool {
	switch(s) {
		case '}':
			return d == '{'
		case ')':
			return d == '('
		case ']':
			return d == '['
	}

	return false
}
