package leetcode

type CustomStack struct {
	Elements []int
	TopIndex int
}

func ConstructorCustomStack(maxSize int) CustomStack {
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
