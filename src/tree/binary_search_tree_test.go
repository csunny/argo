package tree

import (
	"testing"
	"fmt"
)

var bst ItemBinarySearchTree

func fillTree(bst *ItemBinarySearchTree)  {
	bst.Insert(8, "8")
	bst.Insert(3, "3")
	bst.Insert(1, "1")
	bst.Insert(10, "10")
	bst.Insert(6, "6")
	bst.Insert(4, "4")
	bst.Insert(7, "7")
	bst.Insert(13, "13")
	bst.Insert(14, "14")
}

func TestItemBinarySearchTree_InOrderTraverse(t *testing.T) {
	fillTree(&bst)
	bst.String()

	var result []string
	bst.InOrderTraverse(func(i Item) {
		result = append(result, fmt.Sprintf("%s", i))
	})

	t.Error(result)
}

func TestItemBinarySearchTree_PreOrderTraverse(t *testing.T) {

	var result []string
	bst.PostOrderTraverse(func(i Item) {
		result = append(result, fmt.Sprintf("%s", i))
	})

	t.Error(result)
}

func TestItemBinarySearchTree_PostOrderTraverse(t *testing.T) {
	var result []string
	bst.PreOrderTraverse(func(i Item) {
		result = append(result, fmt.Sprintf("%s", i))
	})

	t.Error(result)
}
