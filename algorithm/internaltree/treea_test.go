package internaltree

import (
	"fmt"
	"testing"
)

// https://blog.csdn.net/zearot/article/details/48299459
func TestInternalTree_Find(t *testing.T) {
	tree := InternalTree{}
	tree.BuildTree(1, 15)

	tree.Insert(1, 8)
	tree.Insert(1, 6)
	tree.Insert(1, 5)
	tree.Insert(1, 5)
	tree.Insert(1, 5)
	tree.Insert(1, 5)
	tree.Insert(1, 5)

	tree.Bfs()

	fmt.Println(tree.Find(1,3))
}