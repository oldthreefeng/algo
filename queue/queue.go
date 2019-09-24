/*
@Time : 2019/9/23 20:55
@Author : louis
@File : queue
@Software: GoLand
*/

package queue

type Item struct {
	item interface{}
	prev *Item
}

// Base data structure for Queue
type Queue struct {
	current *Item
	last *Item
	depth uint64
}

// Initializes new Queue and return it
func New() *Queue {
	var queue = new(Queue)

	queue.depth = 0

	return queue
}

// Puts a given item into Queue
func (queue *Queue) Enqueue(item interface{}) {
	if queue.depth == 0 {
		queue.current = &Item{item: item, prev: nil}
		queue.last = queue.current
		queue.depth++
		return
	}

	q := &Item{item: item, prev: nil}
	queue.last.prev = q
	queue.last = q
	queue.depth++
}

// Extracts first item from the Queue
func (queue *Queue) Dequeue() interface{} {
	if !queue.IsEmpty() {
		item := queue.current.item
		queue.current = queue.current.prev
		queue.depth--

		return item
	}

	return nil
}

// IsEmpty
func (queue *Queue) IsEmpty() bool  {
	return queue.depth == 0
}

// Peek first
func (queue *Queue) Peek() interface{}  {
	if queue.IsEmpty() {
		return nil
	}
	return queue.current.item
}

func (queue *Queue) PeekBack() interface{} {
	if queue.IsEmpty() {
		return nil
	}
	return queue.last.item
}