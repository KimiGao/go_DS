// 循环队列的实现

package main

import (
	"fmt"
)

const MAXSIZE = 10

type LoopQueue struct {
	dataArray [MAXSIZE]int
	front     int
	rear      int
}

func (q *LoopQueue) isEmpty() bool {
	return q.front == q.rear
}

func (q *LoopQueue) isFull() bool {
	return ((q.rear + 1) % MAXSIZE) == q.front
}

func (q *LoopQueue) length() int {
	return (q.rear - q.front + MAXSIZE) % MAXSIZE
}

func (q *LoopQueue) insert(data int) int {
	if q.isFull() {
		return -1
	}

	q.dataArray[q.rear] = data
	q.rear = (q.rear + 1) % MAXSIZE

	return q.rear
}

func (q *LoopQueue) remove() int {
	if q.isEmpty() {
		return -1
	}

	q.dataArray[q.front] = 0
	q.front = (q.front + 1) % MAXSIZE

	return q.front
}

func (q *LoopQueue) printInfo() {
	fmt.Printf("queue: %v\nlength: %v\nfront: %v\nrear: %v\n\n", q.dataArray, q.length(), q.front, q.rear)
}

func main() {
	queue := new(LoopQueue)

	queue.insert(1)
	queue.insert(2)
	queue.insert(3)
	queue.insert(4)
	queue.insert(5)
	queue.insert(6)
	queue.insert(7)
	queue.printInfo()

	queue.remove()
	queue.remove()
	queue.printInfo()

	queue.insert(8)
	queue.insert(9)
	queue.insert(10)
	queue.insert(11)
	queue.insert(12)
	queue.insert(13)
	queue.insert(14)
	queue.printInfo()

}
