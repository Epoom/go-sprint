package sprint

func BalancedParentheses(input string) bool {
	stack := []rune{}

	parentheses := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}

	for _, c := range input {
		switch c {
			case '(', '{', '[':
				stack = append(stack, c)
			case ')', '}', ']':
				if len(stack) == 0 || stack[len(stack)-1] != parentheses[c] {
					return false
				}
				stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}
