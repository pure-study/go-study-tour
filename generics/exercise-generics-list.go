package main

import (
	"errors"
	"fmt"
)

// List represents a singly-linked list that holds
// values of any type.
type List[T any] struct {
	next *List[T]
	val  T
}

func main() {
	list := List[int]{nil, 1}
	list.Append(List[int]{nil, 2}).Append(List[int]{nil, 3}).Append(List[int]{nil, 4})
	list.InsertSingle(&List[int]{nil, 5})
	fmt.Println(list)

	list.DeleteAt(1)
	fmt.Println(list)
	_, err := list.DeleteAt(5)
	fmt.Println(err)
	fmt.Println(list)

	list.DeleteAt(0)
	fmt.Println(list)
}

// always append to the last, and return the original list itself
func (lst *List[Type]) Append(n List[Type]) *List[Type] {
	last := lst
	for last.next != nil {
		last = last.next
	}
	last.next = &n
	return lst
}

// insert to the next of current position, and return the original list itself
func (lst *List[Type]) InsertSingle(n *List[Type]) *List[Type] {
	if n == nil || n.next != nil {
		panic("unsupported operation for now!")
	}

	currentNext := lst.next
	lst.next = n
	n.next = currentNext
	return lst
}

// index starts from 0
func (lst *List[Type]) DeleteAt(index int) (*List[Type], error) {
	p, c, n := (*List[Type])(nil), lst, lst.next
	for i := 0; i < index; i++ {
		if c.next == nil {
			return lst, errors.New(fmt.Sprintf("the element at %d not found", index))
		}
		p, c, n = c, n, n.next
	}

	if p == nil {
		lst.val = n.val
		lst.next = n.next
	} else {
		p.next = n
	}
	return lst, nil
}

func (lst List[Type]) String() string {
	return fmt.Sprintf("%v->%v", lst.val, lst.next)
}
