package domain

type Stack struct {
	content []ArrangementContainer
}

func NewStack() *Stack {
	return &Stack{make([]ArrangementContainer, 0)}
}

func (stack *Stack) Size() int {
	return len(stack.content)
}

func (stack *Stack) IsEmpty() bool {
	return stack.Size() == 0
}

func (stack *Stack) Peek() ArrangementContainer {
	return stack.content[len(stack.content)-1]
}

func (stack *Stack) Pop() ArrangementContainer {
	index := len(stack.content) - 1
	element := stack.content[index]
	stack.content = stack.content[:index]
	return element
}

func (stack *Stack) Push(newItem ArrangementContainer) {
	stack.content = append(stack.content, newItem)
}
