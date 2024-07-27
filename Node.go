package priority

import "cmp"

type Node[K cmp.Ordered, V any] struct {
	Key   K
	Value V
}
