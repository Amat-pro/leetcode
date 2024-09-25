package stack_queue

import "strconv"

// evalRPN 逆波兰表达式求值
func evalRPN(tokens []string) int {
	stack := []int{}

	for _, token := range tokens {
		switch token {
		case "+":
			n1 := stack[len(stack)-2]
			n2 := stack[len(stack)-1]
			stack = stack[:len(stack)-2]
			stack = append(stack, n1+n2)
		case "-":
			n1 := stack[len(stack)-2]
			n2 := stack[len(stack)-1]
			stack = stack[:len(stack)-2]
			stack = append(stack, n1-n2)
		case "*":
			n1 := stack[len(stack)-2]
			n2 := stack[len(stack)-1]
			stack = stack[:len(stack)-2]
			stack = append(stack, n1*n2)
		case "/":
			n1 := stack[len(stack)-2]
			n2 := stack[len(stack)-1]
			stack = stack[:len(stack)-2]
			stack = append(stack, n1/n2)
		default:
			num, _ := strconv.Atoi(token)
			stack = append(stack, num)
		}
	}

	return stack[0]
}