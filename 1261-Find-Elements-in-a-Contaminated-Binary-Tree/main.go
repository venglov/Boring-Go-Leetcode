package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type FindElements struct {
	root    *TreeNode
	visited map[int]bool
}

func Constructor(root *TreeNode) FindElements {
	root.Val = 0
	visited := make(map[int]bool, 10)
	visited[0] = true
	dfs(root, visited)
	return FindElements{root: root, visited: visited}
}

func dfs(tn *TreeNode, visited map[int]bool) {
	if tn.Left != nil {
		tn.Left.Val = tn.Val*2 + 1
		visited[tn.Left.Val] = true
		dfs(tn.Left, visited)
	}
	if tn.Right != nil {
		tn.Right.Val = tn.Val*2 + 2
		visited[tn.Right.Val] = true
		dfs(tn.Right, visited)
	}
}

func dfs_search(tn *TreeNode, target int, found *bool) {
	if tn.Val == target {
		*found = true
		return
	}
	if tn.Left != nil && !*found {
		dfs_search(tn.Left, target, found)
	}
	if tn.Right != nil && !*found {
		dfs_search(tn.Right, target, found)
	}
}

func (this *FindElements) Find(target int) bool {
	// found := false
	// dfs_search(this.root, target, &found)
	return this.visited[target]
}

func main() {
	root := &TreeNode{Val: -1}
	root.Left = &TreeNode{Val: -1}
	root.Right = &TreeNode{Val: -1}
	root.Left.Left = &TreeNode{Val: -1}
	root.Left.Right = &TreeNode{Val: -1}
	root.Right.Right = &TreeNode{Val: -1}

	obj := Constructor(root)
	fmt.Println(obj.Find(1))
	fmt.Println(obj.Find(2))
	fmt.Println(obj.Find(3))
	fmt.Println(obj.Find(4))
	fmt.Println(obj.Find(5))
	fmt.Println(obj.Find(6))
}
