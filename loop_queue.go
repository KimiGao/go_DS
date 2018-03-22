// 循环队列的实现

package ds

import (
	"fmt"
)

const MAXSIZE = 10

type LoopQueue struct {
	dataArray [MAXSIZE]int
	front     int
	rear      int
}

func (q *LoopQueue) IsEmpty() bool {
	return q.front == q.rear
}

func (q *LoopQueue) IsFull() bool {
	return ((q.rear + 1) % MAXSIZE) == q.front
}

func (q *LoopQueue) Length() int {
	return (q.rear - q.front + MAXSIZE) % MAXSIZE
}

func (q *LoopQueue) Insert(data int) int {
	if q.IsFull() {
		return -1
	}

	q.dataArray[q.rear] = data
	q.rear = (q.rear + 1) % MAXSIZE

	return q.rear
}

func (q *LoopQueue) Remove() int {
	if q.IsEmpty() {
		return -1
	}

	q.dataArray[q.front] = 0
	q.front = (q.front + 1) % MAXSIZE

	return q.front
}

func (q *LoopQueue) PrintInfo() {
	fmt.Printf("queue: %v\nlength: %v\nfront: %v\nrear: %v\n\n", q.dataArray, q.Length(), q.front, q.rear)
}
