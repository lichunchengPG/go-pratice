package main

import "fmt"

type Queue []int


// Pushes element into queue
// e.g q.push(123)
func (q *Queue) Push(v int)  {
	*q = append(*q, v)
}

// Pops element from head
func (q *Queue) Pop() int  {
	head := (*q)[0]
	*q = (*q)[1:]
	return head
}

// Return if the queue is empty or not
func (q *Queue) IsEmpty() bool  {
	return len(*q) == 0
}

func main()  {
	q := Queue{1, 2}
	q.Push(3)
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Println(q.IsEmpty())
	fmt.Println(q.Pop())
	fmt.Println(q.IsEmpty())

}