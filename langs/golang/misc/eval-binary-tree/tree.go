package binary_tree

import "golang.org/x/exp/constraints"

type treeNode[T constraints.Ordered] struct {
	val   T
	left  *treeNode[T]
	right *treeNode[T]
}

func newTree[T constraints.Ordered](vv []T) *treeNode[T] {
	if len(vv) == 0 {
		return nil
	}
	root := &treeNode[T]{val: vv[0]}
	for i := 1; i < len(vv); i++ {
		root.insert(vv[i])
	}
	return root
}

func (n *treeNode[T]) insert(val T) { // could include data and then this would take a map instead
	if n == nil {
		return
	}
	if val < n.val {
		if n.left == nil {
			n.left = &treeNode[T]{val: val}
		} else {
			n.left.insert(val)
		}
	} else {
		if n.right == nil {
			n.right = &treeNode[T]{val: val}
		} else {
			n.right.insert(val)
		}
	}
}
