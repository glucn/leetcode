package solution

// Runtime 16 ms
// Memory 6.4 MB

type MyCircularDeque struct {
	queue []int
	size  int
}

/** Initialize your data structure here. Set the size of the deque to be k. */
func Constructor(k int) MyCircularDeque {
	return MyCircularDeque{
		queue: []int{},
		size:  k,
	}
}

/** Adds an item at the front of Deque. Return true if the operation is successful. */
func (q *MyCircularDeque) InsertFront(value int) bool {
	if q.IsFull() {
		return false
	}
	q.queue = append([]int{value}, q.queue...)
	return true
}

/** Adds an item at the rear of Deque. Return true if the operation is successful. */
func (q *MyCircularDeque) InsertLast(value int) bool {
	if q.IsFull() {
		return false
	}
	q.queue = append(q.queue, value)
	return true
}

/** Deletes an item from the front of Deque. Return true if the operation is successful. */
func (q *MyCircularDeque) DeleteFront() bool {
	if q.IsEmpty() {
		return false
	}
	q.queue = q.queue[1:]
	return true
}

/** Deletes an item from the rear of Deque. Return true if the operation is successful. */
func (q *MyCircularDeque) DeleteLast() bool {
	if q.IsEmpty() {
		return false
	}
	q.queue = q.queue[:len(q.queue)-1]
	return true
}

/** Get the front item from the deque. */
func (q *MyCircularDeque) GetFront() int {
	if q.IsEmpty() {
		return -1
	}
	return q.queue[0]
}

/** Get the last item from the deque. */
func (q *MyCircularDeque) GetRear() int {
	if q.IsEmpty() {
		return -1
	}
	return q.queue[len(q.queue)-1]
}

/** Checks whether the circular deque is empty or not. */
func (q *MyCircularDeque) IsEmpty() bool {
	return len(q.queue) == 0
}

/** Checks whether the circular deque is full or not. */
func (q *MyCircularDeque) IsFull() bool {
	return len(q.queue) == q.size
}

/**
 * Your MyCircularDeque object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.InsertFront(value);
 * param_2 := obj.InsertLast(value);
 * param_3 := obj.DeleteFront();
 * param_4 := obj.DeleteLast();
 * param_5 := obj.GetFront();
 * param_6 := obj.GetRear();
 * param_7 := obj.IsEmpty();
 * param_8 := obj.IsFull();
 */
