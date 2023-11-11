package design_a_stack_with_increment_operation

type CustomStack struct {
	Elements []int
	TopIndex int
}

func min(a, b int) int {
	if a <= b {
		return a
	}

	return b
}

func Constructor(maxSize int) CustomStack {
	return CustomStack{make([]int, maxSize), -1}
}

func (stack *CustomStack) Push(x int) {
	if stack.TopIndex+1 == len(stack.Elements) {
		return
	}

	stack.TopIndex++
	stack.Elements[stack.TopIndex] = x
}

func (stack *CustomStack) Pop() int {
	if stack.TopIndex == -1 {
		return -1
	}

	stack.TopIndex--
	return stack.Elements[stack.TopIndex+1]
}

func (stack *CustomStack) Increment(k int, val int) {
	for i := 0; i < min(stack.TopIndex+1, k); i++ {
		stack.Elements[i] += val
	}
}
