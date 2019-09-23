/*
@Time : 2019/9/23 23:21
@Author : louis
@File : queue_test
@Software: GoLand
*/

package queue

import "testing"

func TestQueue(t *testing.T) {
	var queue = New()

	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)
	queue.Enqueue(4)
	queue.Enqueue(5)

	for i := 1; i < 6; i++ {
		item := queue.Dequeue()

		if item != i {
			t.Error("TestQueue failed...", i)
		}
	}

}