package queue

type Queue[T any] interface {
	PushBack(val T)
	PushFront(val T)
	PopFront() (T,bool)
	PopBack() (T,bool)
}

type node[T any] struct {
	Val T
	Prev *node[T]
	Next *node[T]

}

type queue[T any] struct {
	Head *node[T]
	Tail *node[T]
	Size int
}

func NewQueue[T any]() *queue[T] {
	return  &queue[T]{}
}


func (q *queue[T]) PushBack(val T) {
	newNode := &node[T]{Val: val}
	if q.Head == nil{
		q.Head = newNode
		q.Tail = newNode
		q.Head.Prev = nil
	}else{
		newNode.Prev = q.Tail
		q.Tail.Next = newNode
		q.Tail = newNode
	}
		q.Size++

}

func (q *queue[T]) PushFront(val T) {
		newNode := &node[T]{Val: val}
	if q.Head == nil{
		q.Head = newNode
		q.Tail = newNode
		q.Head.Prev = nil
	}else{
	newNode.Next = q.Head
	q.Head.Prev = newNode
	q.Head = newNode}
	q.Size++
}

func (q *queue[T]) PopBack() (T,bool) {
	if q.Head == nil{
		var defaultValue T
		return  defaultValue,false
	}

	value := q.Tail.Val
	q.Tail = q.Tail.Prev
	if q.Tail != nil{
		q.Tail.Next = nil
	}else{
		q.Head = nil
	}
	q.Size--
	return value,true
}

func (q *queue[T]) PopFront() (T,bool) {
	if q.Head == nil {
		var defaultValue T
		return defaultValue,false
	}
	value := q.Head.Val
	q.Head = q.Head.Next
	if q.Head != nil {
		q.Head.Prev = nil
	}else{
		q.Tail = nil
	}
	q.Size--
	return value,true
}