package priority

import (
	"container/heap"
)

// type Queue[K cmp.Ordered, V any] struct{ nodes nodes[K, V] }

// func (q *Queue[K, V]) Len() int { return q.nodes.Len() }
// func (q *Queue[K, V]) Push(priority K, value V) *Node[K, V] {
// 	var node = &Node[K, V]{Key: priority, Value: value}
// 	heap.Push(&q.nodes, node)
// 	return node
// }
// func (q *Queue[K, V]) Pop() *Node[K, V] { return q.Remove(0) }
// func (q *Queue[K, V]) Remove(index int) *Node[K, V] {
// 	return heap.Remove(&q.nodes, index).(*Node[K, V])
// }

// type nodes[K cmp.Ordered, V any] []*Node[K, V]

// func (n nodes[K, V]) Len() int           { return len(n) }
// func (n nodes[K, V]) Less(i, j int) bool { return n[i].Key < n[j].Key }
// func (n nodes[K, V]) Swap(i, j int)      { n[i], n[j] = n[j], n[i] }
// func (n *nodes[K, V]) Push(x any)        { *n = append(*n, x.(*Node[K, V])) }
//
//	func (n *nodes[K, V]) Pop() any {
//		var node = (*n)[len(*n)-1]
//		*n = (*n)[0 : len(*n)-1]
//		return node
//	}
type Queue[V Ordered] []V

func (q *Queue[V]) Len() int             { return len(*q) }
func (q *Queue[V]) Less(i, j int) bool   { return (*q)[i].Less((*q)[j]) }
func (q *Queue[V]) Swap(i, j int)        { (*q)[i], (*q)[j] = (*q)[j], (*q)[i] }
func (q *Queue[V]) Push(v any)           { heap.Push(q, v) }
func (q *Queue[V]) Pop() any             { return q.Remove(0) }
func (q *Queue[V]) Remove(index int) any { return heap.Remove(q, index) }
