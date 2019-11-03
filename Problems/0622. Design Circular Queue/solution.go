package solution

// Runtime 16 ms
// Memory 6.2 MB

type MyCircularQueue struct {
	head     *node
	tail     *node
	capacity int
	length   int
}

type node struct {
	value int
	next  *node
}

/** Initialize your data structure here. Set the size of the queue to be k. */
func Constructor(k int) MyCircularQueue {
	return MyCircularQueue{
		capacity: k,
	}
}

/** Insert an element into the circular queue. Return true if the operation is successful. */
func (q *MyCircularQueue) EnQueue(value int) bool {
	if q.IsFull() {
		return false
	}

	if q.tail == nil {
		q.head = &node{
			value: value,
		}
		q.tail = q.head
		q.length++
		return true
	}
	q.tail.next = &node{
		value: value,
	}
	q.tail = q.tail.next
	q.length++
	return true
}

/** Delete an element from the circular queue. Return true if the operation is successful. */
func (q *MyCircularQueue) DeQueue() bool {
	if q.IsEmpty() {
		return false
	}
	if q.tail == q.head {
		q.tail, q.head = nil, nil
		q.length--
		return true
	}
	q.head = q.head.next
	q.length--
	return true
}

/** Get the front item from the queue. */
func (q *MyCircularQueue) Front() int {
	if q.IsEmpty() {
		return -1
	}
	return q.head.value
}

/** Get the last item from the queue. */
func (q *MyCircularQueue) Rear() int {
	if q.IsEmpty() {
		return -1
	}
	return q.tail.value
}

/** Checks whether the circular queue is empty or not. */
func (q *MyCircularQueue) IsEmpty() bool {
	return q.length == 0
}

/** Checks whether the circular queue is full or not. */
func (q *MyCircularQueue) IsFull() bool {
	return q.length == q.capacity
}

/**
 * Your MyCircularQueue object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.EnQueue(value);
 * param_2 := obj.DeQueue();
 * param_3 := obj.Front();
 * param_4 := obj.Rear();
 * param_5 := obj.IsEmpty();
 * param_6 := obj.IsFull();
 */
