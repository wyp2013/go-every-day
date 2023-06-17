package internaltree

import (
	"container/list"
	"fmt"
)

type InternalTree struct {
	root *treeNode
}

type treeNode struct {
	left        int
	right       int
	cover       int
	leftChild   *treeNode
	rightChild  *treeNode
}

func newTreeNode(a int, b int) *treeNode {
	return &treeNode{
		left: a,
		right: b,
		cover: 0,
		leftChild: nil,
		rightChild: nil,
	}
}

func (t *InternalTree) BuildTree(a int, b int) {
	t.root = newTreeNode(a, b)
	t.root.buildTree(a, b)
}

func (t *InternalTree) check() {
	if t.root == nil {
		panic("root is empty")
	}
}

func (t *InternalTree) Insert(a int, b int) {
	t.check()
	t.root.insert(a, b)
}

func (t *InternalTree) Delete(a int, b int) {
	t.check()
	t.root.delete(a, b)
}

func (t *InternalTree) Find(a int, b int) int {
	t.check()
	return t.root.find(a, b)
}

func (t *InternalTree) Bfs() {
	t.check()
	t.root.BreadthFirstSearch()
}

func (n *treeNode) buildTree(a int, b int) {
	if b > a {
		mid := (a+b)/2
		n.leftChild  = newTreeNode(a, mid)
		n.leftChild.buildTree(a, mid)

		n.rightChild = newTreeNode(mid+1, b)
		n.rightChild.buildTree(mid+1, b)
	}
}

func (n *treeNode) insert(a int, b int) {
	if n.left == a && n.right == b {
		n.cover ++
		return
	}

	mid := (n.left + n.right)/2
	if a > mid {
		n.rightChild.insert(a, b)
	} else if b <= mid {
		n.leftChild.insert(a, b)
	} else {
		n.leftChild.insert(a, mid)
		n.rightChild.insert(mid+1, b)
	}
}

func (n *treeNode) delete(a int, b int) {
	if n.left == a && n.right == b {
		if n.cover > 0 {
			n.cover--
		}
		return
	}

	mid := (n.left + n.right)/2
	if a > mid {
		n.rightChild.delete(a, b)
	} else if b <= mid {
		n.leftChild.delete(a, b)
	} else {
		n.leftChild.delete(a, mid)
		n.rightChild.delete(mid+1, b)
	}
}

func (n *treeNode) find(a int, b int) int {
	if n.left == a && n.right == b {
		return n.cover
	}

	mid := (n.left + n.right)/2
	if a > mid {
		return n.rightChild.find(a, b)
	} else if b <= mid {
		return n.leftChild.find(a, b)
	} else {
		leftCover  := n.leftChild.find(a, mid)
		rightCover := n.rightChild.find(mid+1, b)

		if leftCover < rightCover {
			leftCover = rightCover
		}
		return n.cover + leftCover
	}
}

func (n *treeNode) String() string {
	return fmt.Sprintf("[left: %d right: %d cover:%d]", n.left, n.right, n.cover)
}

func (n *treeNode) BreadthFirstSearch() {
	que := list.New()
	que.PushBack(n)

	for que.Len() > 0 {
		e := que.Front()

		n := e.Value.(*treeNode)
		if n.cover > 0 {
			fmt.Println(n)
		}
		que.Remove(e)

		if n.leftChild != nil {
			que.PushBack(n.leftChild)
		}
		if n.rightChild != nil {
			que.PushBack(n.rightChild)
		}
	}
}



