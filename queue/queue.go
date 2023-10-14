package main

import (
	"errors"
	"fmt"
)

const LEN = 50

type Queue struct {
	data [LEN]int
	head int
	tail int
	size int
}

func enqueue(q *Queue, value int) error {
	if q.size == LEN {
		return errors.New("Full Queue")
	}

	if q.size == 0 {
		q.data[0] = value
		q.tail++
		q.size++

		return nil
	}

	q.data[q.tail] = value
	q.tail++
	q.size++

	return nil
}

func dequeue(q *Queue) (int, error) {
	if q.size == 0 {
		return 0, errors.New("Empty Queue")
	}

	trash := q.data[q.head]
	for i := 0; i < q.size - 1; i++ {
		q.data[i] = q.data[i + 1]
	}

	q.tail--
	q.size--

	return trash, nil
}

func show(q *Queue) {
	for _, data := range q.data {
		if data != 0 {
			fmt.Printf("%d ", data)
		}
	}

	fmt.Println()
}

func init_queue() *Queue {
	var data [50]int

	q := &Queue{head: 0, tail: 0, size: 0, data: data}
	return q
} 

func main() {
	queue := init_queue()

	enqueue(queue, 12)
	enqueue(queue, 10)
	enqueue(queue, 3)
	enqueue(queue, 5)
	enqueue(queue, 6)
	enqueue(queue, 8)
	
	dequeue(queue)
	dequeue(queue)

	show(queue)
}
