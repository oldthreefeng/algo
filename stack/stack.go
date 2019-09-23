/*
@Time : 2019/9/23 23:29
@Author : louis
@File : stack
@Software: GoLand
*/

package stack

type Item struct {
	item interface{}
	next *Item
}

// Stack is a base structure for LIFO
type Stack struct {
	sp    *Item
	depth uint64
}

// Initialzes new Stack
func New() *Stack {
	var stack = new(Stack)

	stack.depth = 0
	return stack
}

// Pushes a given item into Stack
func (stack *Stack) Push(item interface{}) {
	stack.sp = &Item{item: item, next: stack.sp}
	stack.depth++
}

// Deletes top of a stack and return it
func (stack *Stack) Pop() interface{} {
	if stack.depth > 0 {
		item := stack.sp.item
		stack.sp = stack.sp.next
		stack.depth--
		return item
	}

	return nil
}

// Peek returns top of a stack without deletion
func (stack *Stack) Peek() interface{} {
	if stack.depth > 0 {
		return stack.sp.item
	}

	return nil
}

func (stack *Stack) IsEmpty() bool  {
	return stack.depth == 0
}


