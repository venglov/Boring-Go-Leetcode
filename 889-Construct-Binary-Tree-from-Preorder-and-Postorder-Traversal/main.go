package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func constructFromPrePost(preorder []int, postorder []int) *TreeNode {
	root := &TreeNode{}
	dfs(root, 0, 0, &preorder, &postorder)
	return root
}

func dfs(tn *TreeNode, preIndex int, postIndex int, preorder *[]int, postorder *[]int) (int, int) {
	if preIndex >= len(*preorder) || postIndex >= len(*postorder) {
		return postIndex, preIndex
	}
	if (*preorder)[preIndex] == (*postorder)[postIndex] {
		tn.Val = (*preorder)[preIndex]
		return postIndex + 1, preIndex
	}
	newPreIndex := preIndex

	lChild := &TreeNode{}
	tn.Left = lChild
	postIndex, newPreIndex = dfs(lChild, newPreIndex+1, postIndex, preorder, postorder)

	if (*preorder)[preIndex] == (*postorder)[postIndex] {
		tn.Val = (*preorder)[preIndex]
		return postIndex + 1, newPreIndex
	}

	rChild := &TreeNode{}
	tn.Right = rChild
	postIndex, newPreIndex = dfs(rChild, newPreIndex+1, postIndex, preorder, postorder)

	if (*preorder)[preIndex] == (*postorder)[postIndex] {
		tn.Val = (*preorder)[preIndex]
		return postIndex + 1, newPreIndex
	}

	return postIndex + 1, preIndex
}

func main() {
	preorder := []int{1, 2, 4, 5, 3, 6, 7}
	postorder := []int{4, 5, 2, 6, 7, 3, 1}

	root := constructFromPrePost(preorder, postorder)
	fmt.Println(root)
}
