package ds

import "errors"

type Comparable interface {
	comparable
}

type List[T Comparable] struct {
	head   *node[T]
	tail   *node[T]
	Length int
}

type node[T Comparable] struct {
	Value T
	Next  *node[T]
	Prev  *node[T]
}

func (lst *List[T]) RemoveAt(idx int) (T, error) {
	if idx > lst.Length || idx < 0 {
		return *new(T), errors.New("bad index")
	}
	curr := lst.head
	for i := 0; curr != nil && i < idx; i++ {
		curr = curr.Next
	}

	lst.Length--

	return lst.removeNode(curr), nil
}

func (lst *List[T]) Remove(value T) (T, error) {
	curr := lst.head
	lst.Length--
	if lst.Length == 0 {

		lst.head = nil
		lst.tail = nil
		return curr.Value, nil

	} else if lst.Length < 0 {
		lst.Length = 0
		return *new(T), errors.New("empty list brother man")
	}

	for curr != nil {
		if curr.Value == value {

			return lst.removeNode(curr), nil
		}

		curr = curr.Next
	}

	return *new(T), errors.New("value not found")
}

func (lst *List[T]) Append(value T) {
	node := node[T]{Value: value}
	lst.Length++
	if lst.tail == nil {
		lst.tail = &node
		lst.head = &node
		return
	}

	node.Prev = lst.tail
	lst.tail.Next = &node

	lst.tail = &node

}

func (lst *List[T]) Prepend(value T) {
	node := node[T]{Value: value}
	lst.Length++
	if lst.head == nil {
		lst.tail = &node
		lst.head = &node
		return
	}

	node.Next = lst.head
	lst.head.Prev = &node

	lst.head = &node
}

func (lst *List[T]) InstertAt(idx int, value T) error {
	node := node[T]{Value: value}

	if lst.tail == nil || lst.head == nil {
		return errors.New("empty list my guy")
	}

	if idx > lst.Length {
		return errors.New("oh no")
	}

	if idx < 0 {
		return errors.New("que haces bro")
	}

	if idx == 0 {
		lst.Prepend(value)
		return nil
	} else if idx == lst.Length {
		lst.Append(value)
		return nil
	}
	lst.Length++

	curr := lst.head
	curr = lst.getAt(curr, idx)

	node.Next = curr
	node.Prev = curr.Prev

	curr.Prev.Next = &node
	curr.Prev = &node

	return nil
}

func (lst *List[T]) Get(idx int) T {
	curr := lst.head

	if node := lst.getAt(curr, idx); node != nil {
		return node.Value
	}

	return *new(T)
}

func (lst *List[T]) GetAll() []T {
	var elems []T
	if lst.Length == 0 {
		return elems
	}

	for curr := lst.head; curr != nil; curr = curr.Next {
		elems = append(elems, curr.Value)
	}
	return elems
}

func (lst *List[T]) Clear() {
	lst.head = nil
	lst.tail = nil
	lst.Length = 0
}

func (lst *List[T]) IsEmpty() bool {
	return lst.Length == 0
}

func (lst *List[T]) Peek() (T, T) {

	if lst.head == nil || lst.tail == nil {
		return *new(T), *new(T)
	}

	return lst.head.Value, lst.tail.Value
}

func (lst *List[T]) PeekBack() T {

	if lst.tail == nil {
		return *new(T)
	}

	return lst.tail.Value
}

func (lst *List[T]) PeekFront() T {

	if lst.head == nil {
		return *new(T)
	}

	return lst.head.Value
}

func (lst *List[T]) removeNode(node *node[T]) T {
	if node.Next != nil {
		node.Next.Prev = node.Prev
	}

	if node.Prev != nil {
		node.Prev.Next = node.Next
	}

	if node == lst.head {
		lst.head = node.Next
	}
	if node == lst.tail {
		lst.tail = node.Prev
	}

	node.Prev = nil
	node.Next = nil

	return node.Value
}

func (lst *List[T]) getAt(node *node[T], idx int) *node[T] {
	if idx > lst.Length || idx < 0 {
		return nil
	}
	for i := 0; node != nil && i < idx; i++ {
		node = node.Next
	}

	return node
}
