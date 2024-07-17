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
		var empty T
		return empty, errors.New("bad index")
	}
	curr := lst.head
	for i := 0; curr != nil && i < idx; i++ {
		curr = curr.Next
	}

	if curr.Next != nil {
		curr.Next.Prev = curr.Prev
	}

	if curr.Prev != nil {
		curr.Prev.Next = curr.Next
	}

	if curr == lst.head {
		lst.head = curr.Next
	}
	if curr == lst.tail {
		lst.tail = curr.Prev
	}

	curr.Prev = nil
	curr.Next = nil

	return curr.Value, nil
}

func (lst *List[T]) Remove(value T) (T, error) {
	curr := lst.head
	var empty T

	if lst.Length == 0 {
		lst.Length--

		lst.head = nil
		lst.tail = nil
		return empty, nil

	} else if lst.Length < 0 {
		lst.Length = 0
		return empty, errors.New("list empty brother man")
	}

	for curr != nil {
		if curr.Value == value {
			if curr.Next != nil {
				curr.Next.Prev = curr.Prev
			}

			if curr.Prev != nil {
				curr.Prev.Next = curr.Next
			}

			if curr == lst.head {
				lst.head = curr.Next
			}
			if curr == lst.tail {
				lst.tail = curr.Prev
			}

			curr.Prev = nil
			curr.Next = nil

			lst.Length--

			return curr.Value, nil
		}

		curr = curr.Next
	}

	return empty, errors.New("value not found")
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
	for i := 0; curr != nil && i < idx; i++ {
		curr = curr.Next
	}

	node.Next = curr
	node.Prev = curr.Prev

	curr.Prev.Next = &node
	curr.Prev = &node

	return nil
}

func (lst *List[T]) Get(idx int) (T, error) {
	if idx > lst.Length || idx < 0 {
		var empty T
		return empty, errors.New("bad index")
	}
	curr := lst.head
	for i := 0; curr != nil && i < idx; i++ {
		curr = curr.Next
	}

	return curr.Value, nil
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
