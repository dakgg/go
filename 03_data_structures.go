package main

import "fmt"

// --- Stack (슬라이스 기반) ---

// Stack is a generic LIFO data structure.
type Stack[T any] struct {
	items []T
}

// Push adds v to the top of the stack.
func (s *Stack[T]) Push(v T) {
	s.items = append(s.items, v)
}

// Pop removes and returns the top element.
// Returns the zero value and false if the stack is empty.
func (s *Stack[T]) Pop() (T, bool) {
	if len(s.items) == 0 {
		var zero T
		return zero, false
	}
	top := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return top, true
}

// Peek returns the top element without removing it.
func (s *Stack[T]) Peek() (T, bool) {
	if len(s.items) == 0 {
		var zero T
		return zero, false
	}
	return s.items[len(s.items)-1], true
}

// Len returns the number of elements.
func (s *Stack[T]) Len() int { return len(s.items) }

// --- Queue (슬라이스 기반) ---

// Queue is a generic FIFO data structure.
type Queue[T any] struct {
	items []T
}

// Enqueue adds v to the back of the queue.
func (q *Queue[T]) Enqueue(v T) {
	q.items = append(q.items, v)
}

// Dequeue removes and returns the front element.
func (q *Queue[T]) Dequeue() (T, bool) {
	if len(q.items) == 0 {
		var zero T
		return zero, false
	}
	front := q.items[0]
	q.items = q.items[1:]
	return front, true
}

// Len returns the number of elements.
func (q *Queue[T]) Len() int { return len(q.items) }

// --- Linked List (단방향) ---

// node is a singly linked list node.
type node[T any] struct {
	val  T
	next *node[T]
}

// LinkedList is a generic singly linked list.
type LinkedList[T any] struct {
	head *node[T]
	size int
}

// Prepend inserts v at the front in O(1).
func (l *LinkedList[T]) Prepend(v T) {
	l.head = &node[T]{val: v, next: l.head}
	l.size++
}

// Append inserts v at the back in O(n).
func (l *LinkedList[T]) Append(v T) {
	n := &node[T]{val: v}
	if l.head == nil {
		l.head = n
	} else {
		cur := l.head
		for cur.next != nil {
			cur = cur.next
		}
		cur.next = n
	}
	l.size++
}

// ToSlice returns all values as a slice.
func (l *LinkedList[T]) ToSlice() []T {
	result := make([]T, 0, l.size)
	for cur := l.head; cur != nil; cur = cur.next {
		result = append(result, cur.val)
	}
	return result
}

// --- Binary Search Tree ---

// bstNode is a BST node storing int values.
type bstNode struct {
	val         int
	left, right *bstNode
}

// BST is a binary search tree.
type BST struct {
	root *bstNode
}

// Insert adds val into the BST.
func (t *BST) Insert(val int) {
	t.root = bstInsert(t.root, val)
}

func bstInsert(n *bstNode, val int) *bstNode {
	if n == nil {
		return &bstNode{val: val}
	}
	switch {
	case val < n.val:
		n.left = bstInsert(n.left, val)
	case val > n.val:
		n.right = bstInsert(n.right, val)
	}
	return n
}

// Contains reports whether val is in the BST.
func (t *BST) Contains(val int) bool {
	cur := t.root
	for cur != nil {
		switch {
		case val < cur.val:
			cur = cur.left
		case val > cur.val:
			cur = cur.right
		default:
			return true
		}
	}
	return false
}

// InOrder returns values in ascending order (left → root → right).
func (t *BST) InOrder() []int {
	var result []int
	var traverse func(*bstNode)
	traverse = func(n *bstNode) {
		if n == nil {
			return
		}
		traverse(n.left)
		result = append(result, n.val)
		traverse(n.right)
	}
	traverse(t.root)
	return result
}

// --- Set (맵 기반) ---

// Set is a generic set backed by a map.
type Set[T comparable] struct {
	m map[T]struct{}
}

// NewSet returns an empty Set.
func NewSet[T comparable]() *Set[T] {
	return &Set[T]{m: make(map[T]struct{})}
}

// Add inserts v into the set.
func (s *Set[T]) Add(v T) { s.m[v] = struct{}{} }

// Contains reports whether v is in the set.
func (s *Set[T]) Contains(v T) bool {
	_, ok := s.m[v]
	return ok
}

// Remove deletes v from the set.
func (s *Set[T]) Remove(v T) { delete(s.m, v) }

// Len returns the number of elements.
func (s *Set[T]) Len() int { return len(s.m) }

func dataStructures() {
	fmt.Println("\n=== 자료구조 ===")

	// Stack
	fmt.Println("-- Stack --")
	var st Stack[int]
	for _, v := range []int{1, 2, 3} {
		st.Push(v)
	}
	for st.Len() > 0 {
		v, _ := st.Pop()
		fmt.Printf("  pop: %d\n", v) // 3 2 1 (LIFO)
	}

	// Queue
	fmt.Println("-- Queue --")
	var q Queue[string]
	for _, v := range []string{"a", "b", "c"} {
		q.Enqueue(v)
	}
	for q.Len() > 0 {
		v, _ := q.Dequeue()
		fmt.Printf("  dequeue: %s\n", v) // a b c (FIFO)
	}

	// Linked List
	fmt.Println("-- Linked List --")
	var ll LinkedList[int]
	ll.Append(1)
	ll.Append(2)
	ll.Prepend(0)
	fmt.Println("  list:", ll.ToSlice()) // [0 1 2]

	// BST
	fmt.Println("-- BST --")
	var bst BST
	for _, v := range []int{5, 3, 7, 1, 4, 6, 8} {
		bst.Insert(v)
	}
	fmt.Println("  in-order:", bst.InOrder())     // [1 3 4 5 6 7 8]
	fmt.Println("  contains 4:", bst.Contains(4)) // true
	fmt.Println("  contains 9:", bst.Contains(9)) // false

	// Set
	fmt.Println("-- Set --")
	s := NewSet[string]()
	for _, v := range []string{"go", "python", "go", "rust"} {
		s.Add(v)
	}
	fmt.Println("  len:", s.Len())                      // 3 (중복 제거)
	fmt.Println("  contains go:", s.Contains("go"))     // true
	s.Remove("go")
	fmt.Println("  after remove go:", s.Contains("go")) // false
}
