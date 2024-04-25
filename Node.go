package priority

import "cmp"

type Node[P cmp.Ordered, V any] struct {
	Priority P
	Value    V
}
