package priority

import (
	"cmp"
	"container/heap"
)

type Queue[P cmp.Ordered, V any] struct{ nodes nodes[P, V] }

func New[P cmp.Ordered, V any]() *Queue[P, V] { return &Queue[P, V]{} }
func (q *Queue[P, V]) Len() int               { return q.nodes.Len() }
func (q *Queue[P, V]) Push(priority P, value V) *Node[P, V] {
	var node = &Node[P, V]{Priority: priority, Value: value}
	heap.Push(&q.nodes, node)
	return node
}
func (q *Queue[P, V]) Pop() *Node[P, V] { return q.Remove(0) }
func (q *Queue[P, V]) Remove(index int) *Node[P, V] {
	return heap.Remove(&q.nodes, index).(*Node[P, V])
}

type nodes[P cmp.Ordered, V any] []*Node[P, V]

func (n nodes[P, V]) Len() int           { return len(n) }
func (n nodes[P, V]) Less(i, j int) bool { return cmp.Compare(n[i].Priority, n[j].Priority) < 0 }
func (n nodes[P, V]) Swap(i, j int)      { n[i], n[j] = n[j], n[i] }
func (n *nodes[P, V]) Push(x any)        { *n = append(*n, x.(*Node[P, V])) }
func (n *nodes[P, V]) Pop() any {
	var node = (*n)[len(*n)-1]
	*n = (*n)[0 : len(*n)-1]
	return node
}
