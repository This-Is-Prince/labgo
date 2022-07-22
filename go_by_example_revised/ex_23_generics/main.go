package main

import "fmt"

func MapKeys[K comparable, V any](m map[K]V) []K {
	k := make([]K, 0, len(m))
	for key := range m {
		k = append(k, key)
	}
	return k
}

type List[T any] struct {
	head, tail *element[T]
}

type element[T any] struct {
	next *element[T]
	val  T
}

func (lst *List[T]) Push(v T) {
	elem := &element[T]{val: v}
	if lst.tail == nil {
		lst.head = elem
		lst.tail = lst.head
	} else {
		lst.tail.next = elem
		lst.tail = lst.tail.next
	}
}

func (lst *List[T]) GetAll() []T {
	var elems []T
	for e := lst.head; e != nil; e = e.next {
		elems = append(elems, e.val)
	}
	return elems
}

func main() {
	fmt.Println("Generics")

	m := map[string]int{"one": 1, "two": 2, "three": 3}
	fmt.Println(MapKeys(m))

	lst := List[int]{}
	lst.Push(10)
	lst.Push(13)
	lst.Push(23)
	fmt.Println("list:", lst.GetAll())
}
